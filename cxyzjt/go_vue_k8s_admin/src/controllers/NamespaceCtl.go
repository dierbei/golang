package controllers

import (
	"go_vue_k8s_admin/src/services"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type NamespaceCtl struct {
	NamespaceMap *services.NameSpaceMap `inject:"-"`
}

func NewNamespaceCtl() *NamespaceCtl {
	return &NamespaceCtl{}
}

func (ctl *NamespaceCtl) ListAll(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": ctl.NamespaceMap.ListAll(),
	}
}

func (ctl *NamespaceCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/nslist", ctl.ListAll)
}

func (*NamespaceCtl) Name() string {
	return "NsCtl"
}
