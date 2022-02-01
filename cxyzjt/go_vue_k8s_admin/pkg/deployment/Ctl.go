package deployment

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"go_vue_k8s_admin/src/services"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentCtlV2 struct {
	K8sClient *kubernetes.Clientset   `inject:"-"`
	DeployMap *services.DeploymentMap `inject:"-"`
}

func NewDeploymentCtlV2() *DeploymentCtlV2 {
	return &DeploymentCtlV2{}
}

func (ctl *DeploymentCtlV2) LoadDeploy(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	dep, err := ctl.DeployMap.GetDeployment(ns, name) // 原生
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": dep,
	}
}

func (ctl *DeploymentCtlV2) SaveDeployment(c *gin.Context) goft.Json {
	dep := &v1.Deployment{}
	goft.Error(c.ShouldBindJSON(dep))
	_, err := ctl.K8sClient.AppsV1().Deployments(dep.Namespace).Create(c, dep, metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *DeploymentCtlV2) RmDeployment(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	err := ctl.K8sClient.AppsV1().Deployments(ns).Delete(c, name, metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *DeploymentCtlV2) Build(goft *goft.Goft) {
	//路由
	goft.Handle("GET", "/deployments/:ns/:name", ctl.LoadDeploy)
	goft.Handle("POST", "/deployments", ctl.SaveDeployment)
	goft.Handle("DELETE", "/deployments/:ns/:name", ctl.RmDeployment)
}

func (*DeploymentCtlV2) Name() string {
	return "DeploymentCtlV2"
}
