package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mypro/gsrc/pbfiles"
)
func main() {
	client,err:=grpc.DialContext(context.Background(),"grpc.xiaolatiao.cn:30090",grpc.WithInsecure())
	//client,err:=grpc.DialContext(context.Background(),":8080",grpc.WithInsecure())
	if err!=nil{
		log.Fatal(err)
	}
	rsp:=&pbfiles.ProdResponse{}
	err=client.Invoke(context.Background(),
		"/ProdService/GetProd",
		&pbfiles.ProdRequest{ProdId:123},rsp)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(rsp.Result)


}