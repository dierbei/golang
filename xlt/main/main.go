package main

import (
	"awesomeProject/xlt"
	"fmt"
	"log"
)

type myHandler struct {
}

func (h *myHandler) ServerHttp(w xlt.ResponseWriter, r *xlt.Request) {
	fmt.Println("hello xlt")
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
