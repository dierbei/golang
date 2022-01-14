package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	var tlsConfig = &tls.Config{
		InsecureSkipVerify: true, // 忽略证书验证
	}
	var transport http.RoundTripper = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
		DisableCompression:    true,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		server, _ := url.Parse("https://10.0.16.2:6443")
		log.Println("请求路径3" + request.URL.Path)
		log.Println("请求路径3" + request.URL.Path)
		log.Println("请求路径3" + request.URL.Path)
		p := httputil.NewSingleHostReverseProxy(server)
		p.Transport = transport
		p.ServeHTTP(writer, request)
	})

	log.Println("开始反代k8sapi ")
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
