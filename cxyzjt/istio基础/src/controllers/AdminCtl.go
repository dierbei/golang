package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type AdminCtl struct{}
func NewAdminCtl() *AdminCtl{
  return &AdminCtl{}
}
func(*AdminCtl)  Name() string{
	 return "AdminCtl"
}
func(*AdminCtl) AdminInfo(c *gin.Context) goft.Json{
	return gin.H{"message":"这是管理员才能看的数据"}
}
func(this *AdminCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/admin",this.AdminInfo)
}