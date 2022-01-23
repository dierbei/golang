package services

import (
	"fmt"
	"log"

	"go_vue_k8s_admin/src/wscore"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// DeployHandler 处理Deployment回调的Handler
type DeployHandler struct {
	DeployMap  *DeploymentMap     `inject:"-"`
	DepService *DeploymentService `inject:"-"`
}

func (h *DeployHandler) OnAdd(obj interface{}) {
	h.DeployMap.Add(obj.(*v1.Deployment))
	ns := obj.(*v1.Deployment).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type":   "deployments",
			"result": gin.H{"ns": ns, "data": h.DepService.ListAll(ns)},
		},
	)
}

func (h *DeployHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.DeployMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	} else {
		ns := newObj.(*v1.Deployment).Namespace
		wscore.ClientMap.SendAll(
			gin.H{
				"type":   "deployments",
				"result": gin.H{"ns": ns, "data": h.DepService.ListAll(ns)},
			},
		)
	}
}

func (h *DeployHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*v1.Deployment); ok {
		h.DeployMap.Delete(d)
		ns := obj.(*v1.Deployment).Namespace
		wscore.ClientMap.SendAll(
			gin.H{
				"type":   "deployments",
				"result": gin.H{"ns": ns, "data": h.DepService.ListAll(ns)},
			},
		)
	}
}

// PodHandler 处理Pod回调的Handler
type PodHandler struct {
	PodMap     *PodMap     `inject:"-"`
	PodService *PodService `inject:"-"`
}

func (h *PodHandler) OnAdd(obj interface{}) {
	h.PodMap.Add(obj.(*corev1.Pod))
	ns := obj.(*corev1.Pod).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "pods",
			"result": gin.H{"ns": ns,
				"data": h.PodService.PagePods(ns, 1, 5)},
		},
	)
}

func (h *PodHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.PodMap.Update(newObj.(*corev1.Pod))
	if err != nil {
		log.Println(err)
	} else {
		ns := newObj.(*corev1.Pod).Namespace
		wscore.ClientMap.SendAll(
			gin.H{
				"type": "pods",
				"result": gin.H{"ns": ns,
					"data": h.PodService.PagePods(ns, 1, 5)},
			},
		)
	}
}

func (h *PodHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*corev1.Pod); ok {
		h.PodMap.Delete(d)
		ns := obj.(*corev1.Pod).Namespace
		wscore.ClientMap.SendAll(
			gin.H{
				"type": "pods",
				"result": gin.H{"ns": ns,
					"data": h.PodService.PagePods(ns, 1, 5)},
			},
		)
	}
}

// NameSpaceHandler 处理NameSpace回调的Handler
type NameSpaceHandler struct {
	NamespaceMap *NameSpaceMap `inject:"-"`
}

func (h *NameSpaceHandler) OnAdd(obj interface{}) {
	h.NamespaceMap.Add(obj.(*corev1.Namespace))
}

func (h *NameSpaceHandler) OnUpdate(oldObj, newObj interface{}) {
	h.NamespaceMap.Update(newObj.(*corev1.Namespace))
}

func (h *NameSpaceHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*corev1.Namespace); ok {
		h.NamespaceMap.Delete(d)
	}
}

type EventHandler struct {
	EventMap *EventMap `inject:"-"`
}

func (h *EventHandler) OnAdd(obj interface{}) {
	h.storeData(obj, false)
}

func (h *EventHandler) OnUpdate(oldObj, newObj interface{}) {
	h.storeData(newObj, false)
}

func (h *EventHandler) OnDelete(obj interface{}) {
	h.storeData(obj, true)
}

func (h *EventHandler) storeData(obj interface{}, isdelete bool) {
	if event, ok := obj.(*corev1.Event); ok {
		key := fmt.Sprintf("%s_%s_%s", event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)
		if !isdelete {
			h.EventMap.data.Store(key, event)
		} else {
			h.EventMap.data.Delete(key)
		}
	}
}
