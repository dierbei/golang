package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"log"
	"mypro/src/models"
)

type ProdCtl struct{}
func NewProdCtl() *ProdCtl{
  return &ProdCtl{}
}
func(*ProdCtl)  Name() string{
	 return "ProdCtl"
}
func(*ProdCtl) Detail(c *gin.Context) goft.Json{
	log.Println(c.Request.Header)
	id:=c.Param("id")
	prod:=models.MockProdDetail(id)
	prod.Reviews=models.CallReview(prod.Id)
	return prod
}
func(this *ProdCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/prods/:id",this.Detail)
}