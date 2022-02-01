package services

import (
	"go_vue_k8s_admin/src/helpers"
	"go_vue_k8s_admin/src/models"

	"k8s.io/api/core/v1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type NodeService struct {
	NodeMap *NodeMap             `inject:"-"`
	PodMap  *PodMap              `inject:"-"`
	Metric  *versioned.Clientset `inject:"-"`
}

func NewNodeService() *NodeService {
	return &NodeService{}
}

func (svc *NodeService) ListAllNodes() []*models.NodeModel {
	list := svc.NodeMap.ListAll()
	ret := make([]*models.NodeModel, len(list))
	for i, node := range list {
		nodeUsage := helpers.GetNodeUsage(svc.Metric, node)
		ret[i] = &models.NodeModel{
			Name:       node.Name,
			IP:         node.Status.Addresses[0].Address,
			HostName:   node.Status.Addresses[1].Address,
			CreateTime: node.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Labels:     helpers.FilterLabels(node.Labels),
			Taints:     helpers.FilterTaints(node.Spec.Taints),
			Capacity:   models.NewNodeCapacity(node.Status.Capacity.Cpu().Value(), node.Status.Capacity.Memory().Value(), node.Status.Capacity.Pods().Value()),
			Usage:      models.NewNodeUsage(svc.PodMap.GetNum(node.Name), nodeUsage[0], nodeUsage[1]),
		}
	}
	return ret
}

func (svc *NodeService) LoadOriginNode(nodeName string) *v1.Node {
	return svc.NodeMap.Get(nodeName)
}

func (svc *NodeService) LoadNode(nodeName string) *models.NodeModel {
	node := svc.NodeMap.Get(nodeName)
	nodeUsage := helpers.GetNodeUsage(svc.Metric, node)
	return &models.NodeModel{
		Name:         node.Name,
		IP:           node.Status.Addresses[0].Address,
		HostName:     node.Status.Addresses[1].Address,
		OriginLabels: node.Labels,
		OriginTaints: node.Spec.Taints,
		Capacity: models.NewNodeCapacity(node.Status.Capacity.Cpu().Value(),
			node.Status.Capacity.Memory().Value(), node.Status.Capacity.Pods().Value()),
		Usage:      models.NewNodeUsage(svc.PodMap.GetNum(node.Name), nodeUsage[0], nodeUsage[1]),
		CreateTime: node.CreationTimestamp.Format("2006-01-02 15:04:05"),
	}
}
