package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type ProdModel struct {
	Id int
	Name string
	Reviews interface{}
}
func MockProdDetail(id string  ) *ProdModel{
	id_int,err:=strconv.Atoi(id)
	if err!=nil{
		id_int=0
	}
	return &ProdModel{Id:id_int,Name:"测试商品"}
}
func CallReview(id int ) []map[string]interface{}{
	url:="localhost:8081" // 本地地址地址
	if os.Getenv("release")!=""{
		url="reviewsvc" // k8s 服务名
	}
	req,err:=http.NewRequest("GET",fmt.Sprintf("http://%s/reviews/%d",url,id),nil)
	if err!=nil{
		panic(err.Error())
	}
	rsp,err:=http.DefaultClient.Do(req)
	if err!=nil{
		panic(err.Error())
	}
	defer rsp.Body.Close()
	b,_:=ioutil.ReadAll(rsp.Body)
	m:=[]map[string]interface{}{}

	err=json.Unmarshal(b,&m)
	if err!=nil{
		panic(err.Error())
	}
	return m
}
