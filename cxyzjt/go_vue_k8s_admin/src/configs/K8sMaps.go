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

//InitDepMap DeploymentMap
func (m *K8sMaps) InitDepMap() *core.DeploymentMap {
	return &core.DeploymentMap{}
}
