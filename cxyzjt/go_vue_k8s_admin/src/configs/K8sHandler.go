package configs

import "go_vue_k8s_admin/src/core"

type K8sHandler struct {
}

func NewK8sHandler() *K8sHandler {
	return &K8sHandler{}
}

func (h *K8sHandler) DeployHandler() *core.DeployHandler {
	return &core.DeployHandler{}
}

func (h *K8sHandler) PodHandler() *core.PodHandler {
	return &core.PodHandler{}
}
