package controllers

import (
	"go_vue_k8s_admin/src/services"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
)

type PodCtl struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
	PodSvc    *services.PodService  `inject:"-"`
	Helper    *services.Helper      `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (ctl *PodCtl) GetAll(ctx *gin.Context) goft.Json {
	ns := ctx.DefaultQuery("ns", "default")
	page := ctx.DefaultQuery("current", "1") //当前页
	size := ctx.DefaultQuery("size", "5")

	return gin.H{
		"code": 20000,
		"data": ctl.PodSvc.PagePods(ns, ctl.Helper.StrToInt(page, 1),
			ctl.Helper.StrToInt(size, 5)),
	}
}

func (ctl *PodCtl) Containers(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	podname := c.DefaultQuery("name", "")
	return gin.H{
		"code": 20000,
		"data": ctl.PodSvc.GetPodContainer(ns, podname),
	}

}

func (ctl *PodCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods", ctl.GetAll)
	goft.Handle("GET", "/pods/containers", ctl.Containers)
}

func (ctl *PodCtl) Name() string {
	return "DeploymentCtl"
}
