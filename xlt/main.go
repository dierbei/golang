package main

import (
	"awesomeProject/xlt"
	"bytes"
	"fmt"
	"io"
	"log"
)

type myHandler struct {
}

func (h *myHandler) ServerHttp(w xlt.ResponseWriter, r *xlt.Request) {
	buff := &bytes.Buffer{}

	fmt.Fprintf(buff, "[query]name=%s\n", r.Query("name"))
	fmt.Fprintf(buff, "[query]token=%s\n", r.Query("token"))
	fmt.Fprintf(buff, "[cookie]foo1=%s\n", r.Query("foo1"))
	fmt.Fprintf(buff, "[cookie]foo2=%s\n", r.Query("foo2"))
	fmt.Fprintf(buff, "[Header]User-Agent=%s\n", r.Header.Get("User-Agent"))
	fmt.Fprintf(buff, "[Header]Proto=%s\n", r.Proto)
	fmt.Fprintf(buff, "[Header]Method=%s\n", r.Method)
	fmt.Fprintf(buff,"[Addr]Addr=%s\n",r.RemoteAddr)
	fmt.Fprintf(buff,"[Request]%+v\n",r)

	//发送响应报文
	io.WriteString(w, "HTTP/1.1 200 OK\r\n")
	io.WriteString(w, fmt.Sprintf("Content-Length: %d\r\n",buff.Len()))
	io.WriteString(w,"\r\n")
	io.Copy(w,buff)
}

func main() {
	srv := xlt.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &myHandler{},
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("listen server failed. err:%v\n", err)
	}
}
