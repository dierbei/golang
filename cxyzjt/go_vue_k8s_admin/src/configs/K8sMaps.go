package configs

import (
	"go_vue_k8s_admin/src/services"

	"k8s.io/client-go/kubernetes"
)

type K8sMaps struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewK8sMaps() *K8sMaps {
	return &K8sMaps{}
}

//InitDeploymentMap 初始化Deployment存储Map
func (m *K8sMaps) InitDeploymentMap() *services.DeploymentMap {
	return &services.DeploymentMap{}
}

//InitPodMap 初始化Pod存储Map
func (m *K8sMaps) InitPodMap() *services.PodMap {
	return &services.PodMap{}
}

//InitNamespaceMap 初始化Namespace存储Map
func (m *K8sMaps) InitNamespaceMap() *services.NameSpaceMap {
	return &services.NameSpaceMap{}
}

//InitEventMap 初始化Event存储Map
func (m *K8sMaps) InitEventMap() *services.EventMap {
	return &services.EventMap{}
}
