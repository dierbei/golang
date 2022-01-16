package deployment

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sapi/lib"
	"net/http"
)

func RegHandlers(r *gin.Engine) {
	r.POST("/update/deployment/scale", incrReplicas)
	r.POST("/core/deployments",ListAllDeployments)
}

func ListAllDeployments(ctx *gin.Context)  {
	ns := ctx.DefaultQuery("namesapce", "default")
	ctx.JSON(http.StatusOK,gin.H{"message":"Ok","result":ListAll(ns)})
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
