package controllers

import (
	"go_vue_k8s_admin/src/services"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretCtl struct {
	SecretMap     *services.SecretMap     `inject:"-"`
	SecretService *services.SecretService `inject:"-"`
	Client        *kubernetes.Clientset   `inject:"-"`
}

func NewSecretCtl() *SecretCtl {
	return &SecretCtl{}
}

func (*SecretCtl) Name() string {
	return "SecretCtl"
}

func (ctl *SecretCtl) RmSecret(ctx *gin.Context) goft.Json {
	ns := ctx.DefaultQuery("ns", "default")
	name := ctx.DefaultQuery("name", "")
	goft.Error(ctl.Client.CoreV1().Secrets(ns).
		Delete(ctx, name, v1.DeleteOptions{}))
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}

func (ctl *SecretCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": ctl.SecretService.ListSecret(ns), //暂时 不分页
	}
}

func (ctl *SecretCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/secrets", ctl.ListAll)
	goft.Handle("DELETE", "/secrets", ctl.RmSecret)

}
