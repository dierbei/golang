package configs

import (
	"go_vue_k8s_admin/src/core"

	"k8s.io/client-go/kubernetes"
)

type K8sMaps struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewK8sMaps() *K8sMaps {
	return &K8sMaps{}
}

//InitDeploymentMap 初始化Deployment存储Map
func (m *K8sMaps) InitDeploymentMap() *core.DeploymentMap {
	return &core.DeploymentMap{}
}

//InitPodMap 初始化Pod存储Map
func (m *K8sMaps) InitPodMap() *core.PodMap {
	return &core.PodMap{}
}
