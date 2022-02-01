package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"log"
	"math/big"
	rd "math/rand"
	"os"
	"time"
)

const CAFile="./certs/ca.crt"
const CAKey="./certs/ca.key"
func main() {
	//解析k8s ca 和key 文件
	caFile, err := ioutil.ReadFile(CAFile)
	if err != nil {
		log.Fatal(err)
	}
	caBlock, _ := pem.Decode(caFile)

	caCert, err := x509.ParseCertificate(caBlock.Bytes) //ca 证书对象
	if err != nil {
		log.Fatal(err)
	}
	//解析私钥
	keyFile, err := ioutil.ReadFile(CAKey)
	if err != nil {
		log.Fatal(err)
	}
	keyBlock, _ := pem.Decode(keyFile)
	caPriKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes) //是要对象
	if err != nil {
		log.Fatal(err)
	}
	//---------- ---------------------  ---------------------------------
	//构建证书模板
	certTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(rd.Int63()), //证书序列号
		Subject: pkix.Name{
			Country:            []string{"CN"},
			//Organization:       []string{"填的话这里可以用作用户组"},
			//OrganizationalUnit: []string{"可填课不填"},
			Province:           []string{"beijing"},
			CommonName:         "xiaolatiao",//CN
			Locality:           []string{"beijing"},
		},
		NotBefore:             time.Now(),//证书有效期开始时间
		NotAfter:              time.Now().AddDate(1, 0, 0),//证书有效期
		BasicConstraintsValid: true, //基本的有效性约束
		IsCA:                  false,   //是否是根证书
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, //证书用途(客户端认证，数据加密)
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		EmailAddresses:        []string{"xiaolatiao@xiaolatiao.cn"},
	}

	//生成公私钥--秘钥对
	priKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	//创建证书 对象
	clientCert, err := x509.CreateCertificate(rand.Reader, certTemplate, caCert, &priKey.PublicKey, caPriKey)
	if err != nil {
		return
	}

	//编码证书文件和私钥文件
	clientCertPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: clientCert,
	}


	clientCertFile, err := os.OpenFile("./certs/xiaolatiao.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	err = pem.Encode(clientCertFile, clientCertPem)
	if err != nil {
		log.Fatal(err)
	}

	buf := x509.MarshalPKCS1PrivateKey(priKey)
	keyPem := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: buf,
	}
	clientKeyFile, _ := os.OpenFile("./certs/xiaolatiao_key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	err = pem.Encode(clientKeyFile, keyPem)
	if err != nil {
		log.Fatal(err)
	}
}
