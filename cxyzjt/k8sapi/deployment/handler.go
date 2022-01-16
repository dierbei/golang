package deployment

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sapi/core"
	"k8sapi/lib"
	"net/http"
)

func RegHandlers(r *gin.Engine) {
	r.POST("/update/deployment/scale", incrReplicas)
	r.POST("/core/deployments", ListAllDeployments)
	r.POST("/core/pods", ListPodsByDeployment)
}

//根据 deployment名称 获取 pod列表
func ListPodsByDeployment(c *gin.Context) {
	ns := c.DefaultQuery("namespace", "default")
	depname := c.DefaultQuery("deployment", "default")
	dep, err := core.DepMap.GetDeployment(ns, depname)
	lib.CheckError(err)
	rslist, err := core.RsMap.ListByNameSpace(ns) // 根据namespace 取到 所有rs
	lib.CheckError(err)
	labels, err := GetRsLableByDeployment_ListWatch(dep, rslist) //根据deployment过滤出 rs，然后直接获取标签
	lib.CheckError(err)
	c.JSON(200, gin.H{"message": "Ok", "result": ListPodsByLabel(ns, labels)})
}

func ListAllDeployments(ctx *gin.Context) {
	ns := ctx.DefaultQuery("namesapce", "default")
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok", "result": ListAll(ns)})
}

type incrRep struct {
	Namespace      string `json:"ns" binding:"required,min=1"`
	DeploymentName string `json:"deployment" binding:"required,min=1"`
	Dec            bool   `json:"dec"`
}

func incrReplicas(ctx *gin.Context) {
	req := &incrRep{}
	lib.CheckError(ctx.ShouldBindJSON(req))
	getopt := v1.GetOptions{}
	scale, err := lib.K8sClient.AppsV1().Deployments(req.Namespace).GetScale(context.Background(), req.DeploymentName, getopt)
	lib.CheckError(err)
	if req.Dec {
		scale.Spec.Replicas--
	} else {
		scale.Spec.Replicas++
	}
	uptopt := v1.UpdateOptions{}
	_, err = lib.K8sClient.AppsV1().Deployments(req.Namespace).UpdateScale(context.Background(), req.DeploymentName, scale, uptopt)
	lib.CheckError(err)
	lib.Success("ok", ctx)
}
