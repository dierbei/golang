package main

import (
	"google.golang.org/grpc"
	"log"
	"mypro/gsrc/pbfiles"
	"mypro/gsrc/service"
	"net"
)

func main() {
	myserver:=grpc.NewServer()
	//创建服务
	pbfiles.RegisterProdServiceServer(myserver,service.NewProdService())
	//监听8080
	lis,_:=net.Listen("tcp",":8080")
	if err:=myserver.Serve(lis);err!=nil {
		log.Fatal(err)
	}

}