package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type UserCtl struct {
}

func NewUserCtl() *UserCtl {
	return &UserCtl{}
}

func (ctl *UserCtl) login(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": gin.H{
			"token": "admin-token",
		},
	}
}

func (ctl *UserCtl) logout(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *UserCtl) info(c *gin.Context) string {
	c.Header("Content-type", "application/json")
	return `{"code":20000,"data":{"roles":["admin"],
		"introduction":"I am a super administrator","avatar":"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif","name":"Super Admin"}}`

}

func (ctl *UserCtl) Build(goft *goft.Goft) {
	goft.Handle("POST", "/vue-admin-template/user/login", ctl.login)
	goft.Handle("POST", "/vue-admin-template/user/logout", ctl.logout)
	goft.Handle("GET", "/vue-admin-template/user/info", ctl.info)
}

func (*UserCtl) Name() string {
	return "UserCtl"
}
