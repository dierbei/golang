package models

import "k8s.io/api/core/v1"

type NodeModel struct {
	Name         string
	IP           string
	HostName     string
	CreateTime   string
	Labels       []string
	Taints       []string
	Capacity     *NodeCapacity
	Usage        *NodeUsage
	OriginLabels map[string]string //原始标签
	OriginTaints []v1.Taint        //原始污点
}

type PostNodeModel struct {
	Name         string
	OriginLabels map[string]string //原始标签 ---->前端 是一个对象
	OriginTaints []v1.Taint        //原始污点
}

type NodeCapacity struct {
	Cpu    int64
	Memory int64
	Pods   int64
}

func NewNodeCapacity(cpu int64, memory int64, pods int64) *NodeCapacity {
	return &NodeCapacity{Cpu: cpu, Memory: memory, Pods: pods}
}

type NodeUsage struct {
	Pods   int
	Cpu    float64
	Memory float64
}

func NewNodeUsage(pods int, cpu float64, memory float64) *NodeUsage {
	return &NodeUsage{Pods: pods, Cpu: cpu, Memory: memory}
}
