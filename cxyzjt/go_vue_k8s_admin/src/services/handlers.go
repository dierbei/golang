package services

import (
	"fmt"
	"log"

	"go_vue_k8s_admin/src/wscore"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1beta1"
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

type IngressHandler struct {
	IngressMap *IngressMap `inject:"-"`
}

func (h *IngressHandler) OnAdd(obj interface{}) {
	h.IngressMap.Add(obj.(*v1beta1.Ingress))
	ns := obj.(*v1beta1.Ingress).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "ingress",
			"result": gin.H{"ns": ns,
				"data": h.IngressMap.ListAll(ns)},
		},
	)
}

func (h *IngressHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.IngressMap.Update(newObj.(*v1beta1.Ingress))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*v1beta1.Ingress).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "ingress",
			"result": gin.H{"ns": ns,
				"data": h.IngressMap.ListAll(ns)},
		},
	)

}

func (h *IngressHandler) OnDelete(obj interface{}) {
	h.IngressMap.Delete(obj.(*v1beta1.Ingress))
	ns := obj.(*v1beta1.Ingress).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "ingress",
			"result": gin.H{"ns": ns,
				"data": h.IngressMap.ListAll(ns)},
		},
	)
}

type ServiceHandler struct {
	SvcMap *ServiceMap `inject:"-"`
}

func (h *ServiceHandler) OnAdd(obj interface{}) {
	h.SvcMap.Add(obj.(*corev1.Service))
	ns := obj.(*corev1.Service).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "service",
			"result": gin.H{"ns": ns,
				"data": h.SvcMap.ListAll(ns)},
		},
	)
}

func (h *ServiceHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.SvcMap.Update(newObj.(*corev1.Service))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*corev1.Service).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "service",
			"result": gin.H{"ns": ns,
				"data": h.SvcMap.ListAll(ns)},
		},
	)
}

func (h *ServiceHandler) OnDelete(obj interface{}) {
	h.SvcMap.Delete(obj.(*corev1.Service))
	ns := obj.(*corev1.Service).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "service",
			"result": gin.H{"ns": ns,
				"data": h.SvcMap.ListAll(ns)},
		},
	)
}

type SecretHandler struct {
	SecretMap *SecretMap     `inject:"-"`
	SecretSvc *SecretService `inject:"-"`
}

func (svc *SecretHandler) OnAdd(obj interface{}) {
	svc.SecretMap.Add(obj.(*corev1.Secret))
	ns := obj.(*corev1.Secret).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "secret",
			"result": gin.H{"ns": ns,
				"data": svc.SecretSvc.ListSecret(ns)},
		},
	)
}

func (svc *SecretHandler) OnUpdate(oldObj, newObj interface{}) {
	err := svc.SecretMap.Update(newObj.(*corev1.Secret))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*corev1.Secret).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "secret",
			"result": gin.H{"ns": ns,
				"data": svc.SecretSvc.ListSecret(ns)},
		},
	)
}

func (svc *SecretHandler) OnDelete(obj interface{}) {
	svc.SecretMap.Delete(obj.(*corev1.Secret))
	ns := obj.(*corev1.Secret).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "secret",
			"result": gin.H{"ns": ns,
				"data": svc.SecretSvc.ListSecret(ns)},
		},
	)
}

type ConfigMapHandler struct {
	ConfigMap        *ConfigMap        `inject:"-"`
	ConfigMapService *ConfigMapService `inject:"-"`
}

func (h *ConfigMapHandler) OnAdd(obj interface{}) {
	h.ConfigMap.Add(obj.(*corev1.ConfigMap))
	ns := obj.(*corev1.ConfigMap).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "cm",
			"result": gin.H{"ns": ns,
				"data": h.ConfigMapService.ListConfigMap(ns)},
		},
	)
}

func (h *ConfigMapHandler) OnUpdate(oldObj, newObj interface{}) {
	//重点： 只要update返回true 才会发送 。否则不发送
	if h.ConfigMap.Update(newObj.(*corev1.ConfigMap)) {
		ns := newObj.(*corev1.ConfigMap).Namespace
		wscore.ClientMap.SendAll(
			gin.H{
				"type": "cm",
				"result": gin.H{"ns": ns,
					"data": h.ConfigMapService.ListConfigMap(ns)},
			},
		)
	}
}

func (h *ConfigMapHandler) OnDelete(obj interface{}) {
	h.ConfigMap.Delete(obj.(*corev1.ConfigMap))
	ns := obj.(*corev1.ConfigMap).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "cm",
			"result": gin.H{"ns": ns,
				"data": h.ConfigMapService.ListConfigMap(ns)},
		},
	)
}
