package controllers

import (
	"go_vue_k8s_admin/src/models"
	"go_vue_k8s_admin/src/services"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NodeCtl struct {
	NodeService *services.NodeService `inject:"-"`
	Client      *kubernetes.Clientset `inject:"-"`
}

func NewNodeCtl() *NodeCtl {
	return &NodeCtl{}
}
func (ctl *NodeCtl) ListAll(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": ctl.NodeService.ListAllNodes(),
	}
}

func (ctl *NodeCtl) LoadDetail(c *gin.Context) goft.Json {
	nodeName := c.Param("node")
	return gin.H{
		"code": 20000,
		"data": ctl.NodeService.LoadNode(nodeName),
	}
}

func (ctl *NodeCtl) SaveNode(c *gin.Context) goft.Json {
	nodeModel := &models.PostNodeModel{}
	_ = c.ShouldBindJSON(nodeModel)
	node := ctl.NodeService.LoadOriginNode(nodeModel.Name) //取出原始node 信息
	if node == nil {
		panic("no such node")
	}
	node.Labels = nodeModel.OriginLabels      //覆盖标签
	node.Spec.Taints = nodeModel.OriginTaints //覆盖 污点
	_, err := ctl.Client.CoreV1().Nodes().Update(c, node, v1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *NodeCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/nodes", ctl.ListAll)
	goft.Handle("GET", "/nodes/:node", ctl.LoadDetail)
	goft.Handle("POST", "/nodes", ctl.SaveNode)
}

func (*NodeCtl) Name() string {
	return "NodeCtl"
}
