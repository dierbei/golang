package controllers

import (
	"go_vue_k8s_admin/src/models"
	"go_vue_k8s_admin/src/services"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type IngressCtl struct {
	IngressSvc *services.IngressService `inject:"-"`
	K8sClient  *kubernetes.Clientset    `inject:"-"`
}

func NewIngressCtl() *IngressCtl {
	return &IngressCtl{}
}

func (*IngressCtl) Name() string {
	return "IngressCtl"
}

func (ctl *IngressCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": ctl.IngressSvc.ListIngress(ns), //暂时 不分页
	}
}

func (ctl *IngressCtl) PostIngress(c *gin.Context) goft.Json {
	postModel := &models.IngressPost{}
	goft.Error(c.BindJSON(postModel))
	goft.Error(ctl.IngressSvc.PostIngress(postModel))
	return gin.H{
		"code": 20000,
		"data": postModel,
	}
}

func (ctl *IngressCtl) RmIngress(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	name := c.DefaultQuery("name", "")
	goft.Error(ctl.K8sClient.NetworkingV1beta1().Ingresses(ns).Delete(c, name, v1.DeleteOptions{}))
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}

func (ctl *IngressCtl) Build(goft *goft.Goft) {
	goft.Handle("POST", "/ingress", ctl.PostIngress)
	goft.Handle("DELETE", "/ingress", ctl.RmIngress)
	goft.Handle("GET", "/ingress", ctl.ListAll)
}
