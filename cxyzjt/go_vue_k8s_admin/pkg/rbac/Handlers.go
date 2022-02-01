package rbac

import (
	corev1 "k8s.io/api/core/v1"
	"log"

	"go_vue_k8s_admin/src/wscore"

	"github.com/gin-gonic/gin"
	rbacv1 "k8s.io/api/rbac/v1"
)

type RoleHandler struct {
	RoleMap     *RoleMap     `inject:"-"`
	RoleService *RoleService `inject:"-"`
}

func (h *RoleHandler) OnAdd(obj interface{}) {
	h.RoleMap.Add(obj.(*rbacv1.Role))
	ns := obj.(*rbacv1.Role).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "role",
			"result": gin.H{"ns": ns,
				"data": h.RoleService.ListRoles(ns)},
		},
	)
}

func (h *RoleHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.RoleMap.Update(newObj.(*rbacv1.Role))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*rbacv1.Role).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "role",
			"result": gin.H{"ns": ns,
				"data": h.RoleService.ListRoles(ns)},
		},
	)
}

func (h *RoleHandler) OnDelete(obj interface{}) {
	h.RoleMap.Delete(obj.(*rbacv1.Role))
	ns := obj.(*rbacv1.Role).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "role",
			"result": gin.H{"ns": ns,
				"data": h.RoleService.ListRoles(ns)},
		},
	)
}

type RoleBindingHandler struct {
	RoleBindingMap *RoleBindingMap `inject:"-"`
	RoleService    *RoleService    `inject:"-"`
}

func (h *RoleBindingHandler) OnAdd(obj interface{}) {
	h.RoleBindingMap.Add(obj.(*rbacv1.RoleBinding))
	ns := obj.(*rbacv1.RoleBinding).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "rolebinding",
			"result": gin.H{"ns": ns,
				"data": h.RoleService.ListRoles(ns)},
		},
	)
}

func (h *RoleBindingHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.RoleBindingMap.Update(newObj.(*rbacv1.RoleBinding))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*rbacv1.RoleBinding).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "rolebinding",
			"result": gin.H{"ns": ns,
				"data": h.RoleService.ListRoles(ns)},
		},
	)
}

func (h *RoleBindingHandler) OnDelete(obj interface{}) {
	h.RoleBindingMap.Delete(obj.(*rbacv1.RoleBinding))
	ns := obj.(*rbacv1.RoleBinding).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "rolebinding",
			"result": gin.H{"ns": ns,
				"data": h.RoleService.ListRoleBindings(ns)},
		},
	)
}

type ServiceAccountHandler struct {
	SaMap     *ServiceAccountMap     `inject:"-"`
	SaService *ServiceAccountService `inject:"-"`
}

func (h *ServiceAccountHandler) OnAdd(obj interface{}) {
	h.SaMap.Add(obj.(*corev1.ServiceAccount))
	ns := obj.(*corev1.ServiceAccount).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "sa",
			"result": gin.H{"ns": ns,
				"data": h.SaService.ListSa(ns)},
		},
	)
}

func (h *ServiceAccountHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.SaMap.Update(newObj.(*corev1.ServiceAccount))
	if err != nil {
		return
	}
	ns := newObj.(*corev1.ServiceAccount).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "sa",
			"result": gin.H{"ns": ns,
				"data": h.SaService.ListSa(ns)},
		},
	)
}

func (h *ServiceAccountHandler) OnDelete(obj interface{}) {
	h.SaMap.Delete(obj.(*corev1.ServiceAccount))
	ns := obj.(*corev1.ServiceAccount).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "sa",
			"result": gin.H{"ns": ns,
				"data": h.SaService.ListSa(ns)},
		},
	)
}

type ClusterRoleHandler struct {
	ClusterRoleMap *ClusterRoleMap `inject:"-"`
	RoleService    *RoleService    `inject:"-"`
}

func (h *ClusterRoleHandler) OnAdd(obj interface{}) {
	h.ClusterRoleMap.Add(obj.(*rbacv1.ClusterRole))
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "clusterrole",
			"result": gin.H{"ns": "clusterrole",
				"data": h.RoleService.ListClusterRoles()},
		},
	)
}

func (h *ClusterRoleHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.ClusterRoleMap.Update(newObj.(*rbacv1.ClusterRole))
	if err != nil {
		return
	}
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "clusterrole",
			"result": gin.H{"ns": "clusterrole",
				"data": h.RoleService.ListClusterRoles()},
		},
	)
}

func (h *ClusterRoleHandler) OnDelete(obj interface{}) {
	h.ClusterRoleMap.Delete(obj.(*rbacv1.ClusterRole))
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "clusterrole",
			"result": gin.H{"ns": "clusterrole",
				"data": h.RoleService.ListClusterRoles()},
		},
	)
}

type ClusterRoleBindingHandler struct {
	ClusterRoleBindingMap *ClusterRoleBindingMap `inject:"-"`
	RoleService           *RoleService           `inject:"-"`
}

func (h *ClusterRoleBindingHandler) OnAdd(obj interface{}) {
	h.ClusterRoleBindingMap.Add(obj.(*rbacv1.ClusterRoleBinding))
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "clusterrolebinding",
			"result": gin.H{"ns": "clusterrolebinding",
				"data": h.RoleService.ListClusterRoleBindings()},
		},
	)
}

func (h *ClusterRoleBindingHandler) OnUpdate(oldObj, newObj interface{}) {
	err := h.ClusterRoleBindingMap.Update(newObj.(*rbacv1.ClusterRoleBinding))
	if err != nil {
		log.Println(err)
		return
	}
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "clusterrolebinding",
			"result": gin.H{"ns": "clusterrolebinding",
				"data": h.RoleService.ListClusterRoleBindings()},
		},
	)
}

func (h *ClusterRoleBindingHandler) OnDelete(obj interface{}) {
	h.ClusterRoleBindingMap.Delete(obj.(*rbacv1.ClusterRoleBinding))
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "clusterrolebinding",
			"result": gin.H{"ns": "clusterrolebinding",
				"data": h.RoleService.ListClusterRoleBindings()},
		},
	)
}
