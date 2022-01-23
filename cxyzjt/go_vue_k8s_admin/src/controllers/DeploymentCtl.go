package controllers

import (
	"go_vue_k8s_admin/src/services"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
)

type DeploymentCtl struct {
	K8sClient *kubernetes.Clientset       `inject:"-"`
	DeploySvc *services.DeploymentService `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (ctl *DeploymentCtl) GetList(ctx *gin.Context) goft.Json {
	list := ctl.DeploySvc.ListAll("default")
	return gin.H{
		"code": 20000,
		"data": list,
	}
}

func (ctl *DeploymentCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/deployments", ctl.GetList)
}

func (ctl *DeploymentCtl) Name() string {
	return "DeploymentCtl"
}
