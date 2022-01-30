package rbac

import (
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
