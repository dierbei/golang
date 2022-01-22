package core

import (
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"log"
)

// DeployHandler 处理Deployment回调的Handler
type DeployHandler struct {
	DeployMap *DeploymentMap `inject:"-"`
}

func (h *DeployHandler) OnAdd(obj interface{}) {
	h.DeployMap.Add(obj.(*v1.Deployment))
}

func (h *DeployHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.DeployMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	}
}

func (h *DeployHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*v1.Deployment); ok {
		h.DeployMap.Delete(d)
	}
}

// PodHandler 处理Pod回调的Handler
type PodHandler struct {
	PodMap *PodMap `inject:"-"`
}

func (h *PodHandler) OnAdd(obj interface{}) {
	h.PodMap.Add(obj.(*corev1.Pod))
}

func (h *PodHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.PodMap.Update(newObj.(*corev1.Pod))
	if err != nil {
		log.Println(err)
	}
}

func (h *PodHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*corev1.Pod); ok {
		h.PodMap.Delete(d)
	}
}
