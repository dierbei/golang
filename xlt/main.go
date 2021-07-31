package main

import (
	"awesomeProject/xlt"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type myHandler struct {
}

func (h *myHandler) ServerHttp(w xlt.ResponseWriter, r *xlt.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	const prefix = "your message:"
	io.WriteString(w, "HTTP/1.1 200 OK\r\n")
	io.WriteString(w, fmt.Sprintf("Content-Length: %d\r\n", len(buf)+len(prefix)))
	io.WriteString(w, "\r\n")
	io.WriteString(w, prefix)
	w.Write(buf)
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
