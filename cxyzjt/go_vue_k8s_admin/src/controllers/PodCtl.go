package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"go_vue_k8s_admin/src/services"
	"k8s.io/client-go/kubernetes"
)

type PodCtl struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
	PodSvc    *services.PodService  `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (ctl *PodCtl) GetAll(ctx *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": ctl.PodSvc.ListByNs("default"),
	}
}

func (ctl *PodCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods", ctl.GetAll)
}

func (ctl *PodCtl) Name() string {
	return "DeploymentCtl"
}
