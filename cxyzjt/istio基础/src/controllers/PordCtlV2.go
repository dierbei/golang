package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"log"
	"mypro/src/models"
)

type ProdV2Ctl struct{}

func NewProdV2Ctl() *ProdV2Ctl {
	return &ProdV2Ctl{}
}
func (*ProdV2Ctl) Name() string {
	return "ProdCtl"
}

func (*ProdV2Ctl) Detail(c *gin.Context) goft.Json {
	log.Println(c.Request.Header)
	id := c.Param("id")
	prod := models.MockProdDetail(id)
	prod.Reviews = models.CallReview(prod.Id)
	return gin.H{
		"version": "v2",
		"result":  prod,
	}
}
func (this *ProdV2Ctl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/prods/:id", this.Detail)
}
