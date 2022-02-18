package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"mypro/src/models"
)

type ReviewCtl struct{}
func NewReviewCtl() *ReviewCtl{
  return &ReviewCtl{}
}
func(*ReviewCtl)  Name() string{
	 return "ReviewCtl"
}
func(*ReviewCtl) Reviews(c *gin.Context) goft.Json{
	return models.MockReivews()
}
func(this *ReviewCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/reviews/:id",this.Reviews)
}
