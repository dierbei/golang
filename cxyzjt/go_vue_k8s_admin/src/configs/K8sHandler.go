package configs

import (
	"go_vue_k8s_admin/src/services"
)

type K8sHandler struct {
}

func NewK8sHandler() *K8sHandler {
	return &K8sHandler{}
}

func (h *K8sHandler) DeployHandler() *services.DeployHandler {
	return &services.DeployHandler{}
}

func (h *K8sHandler) PodHandler() *services.PodHandler {
	return &services.PodHandler{}
}

func (h *K8sHandler) NamespaceHandler() *services.NameSpaceHandler {
	return &services.NameSpaceHandler{}
}

func (h *K8sHandler) EventHandler() *services.EventHandler {
	return &services.EventHandler{}
}

func (h *K8sHandler) IngressHandler() *services.IngressHandler {
	return &services.IngressHandler{}
}

func (h *K8sHandler) ServiceHandler() *services.ServiceHandler {
	return &services.ServiceHandler{}
}

func (h *K8sHandler) SecretHandler() *services.SecretHandler {
	return &services.SecretHandler{}
}
