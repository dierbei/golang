package core

import (
	v1 "k8s.io/api/apps/v1"
	"log"
)

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
