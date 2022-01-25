package controllers

import (
	"go_vue_k8s_admin/src/services"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type SvcCtl struct {
	SvcMap *services.ServiceMap `inject:"-"`
}

func NewSvcCtl() *SvcCtl {
	return &SvcCtl{}
}

func (ctl *SvcCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": ctl.SvcMap.ListAll(ns),
	}
}

func (ctl *SvcCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/svc", ctl.ListAll)
}

func (*SvcCtl) Name() string {
	return "SvcCtl"
}
