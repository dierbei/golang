package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/shenyisyn/goft-gin/goft"
	"io/ioutil"
	"os"
)

func readPem(pemfile string) []byte {
	getPem, err := os.Open(pemfile)
	if err != nil {
		return nil
	}
	defer getPem.Close()
	b, err := ioutil.ReadAll(getPem)
	if err != nil {
		return nil
	}
	return b
}
func pubKeys(iss string) [][]byte {
	ret := make([][]byte, 0)
	dir := fmt.Sprintf("./pubkeys/%s", iss)
	if fi, err := ioutil.ReadDir(dir); err != nil {
		panic(err.Error())
	} else {
		for _, f := range fi {
			if f.IsDir() {
				continue
			}
			if pem := readPem(fmt.Sprintf("%s/%s", dir, f.Name())); pem != nil {
				ret = append(ret, pem)
			}
		}
	}
	return ret
}

type JwksCtl struct{}

func NewJwksCtl() *JwksCtl {
	return &JwksCtl{}
}

func (this *JwksCtl) Jwks(c *gin.Context) goft.Json {
	iss := c.DefaultQuery("iss", "user.xiaolatiao.cn")
	pems := pubKeys(iss)
	keys := make([]jwk.RSAPublicKey, 0)
	for _, pem := range pems {
		key, err := jwk.ParseKey(pem, jwk.WithPEM(true))
		goft.Error(err)
		if pubKey, ok := key.(jwk.RSAPublicKey); ok {
			keys = append(keys, pubKey)
		}
	}
	return gin.H{"keys": keys}
}

func (this *JwksCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/jwks", this.Jwks)
}
func (*JwksCtl) Name() string {
	return "JwksCtl"
}
