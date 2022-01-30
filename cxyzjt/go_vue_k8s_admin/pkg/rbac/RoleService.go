package rbac

import (
	rbacv1 "k8s.io/api/rbac/v1"
)

type RoleService struct {
	RoleMap        *RoleMap        `inject:"-"`
	RoleBindingMap *RoleBindingMap `inject:"-"`
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (svc *RoleService) ListRoles(ns string) []*RoleModel {
	list := svc.RoleMap.ListAll(ns)
	ret := make([]*RoleModel, len(list))
	for i, item := range list {
		ret[i] = &RoleModel{
			Name:       item.Name,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:  item.Namespace,
		}
	}
	return ret
}

func (svc *RoleService) ListRoleBindings(ns string) []*RoleBindingModel {
	list := svc.RoleBindingMap.ListAll(ns)
	ret := make([]*RoleBindingModel, len(list))
	for i, item := range list {
		ret[i] = &RoleBindingModel{
			Name:       item.Name,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:  item.Namespace,
			Subject:    item.Subjects,
			RoleRef:    item.RoleRef,
		}
	}
	return ret
}

func (svc *RoleService) GetRoleBinding(ns, name string) *rbacv1.RoleBinding {
	rb := svc.RoleBindingMap.Get(ns, name)
	if rb == nil {
		panic("no such rolebinding")
	}
	return rb
}

func (svc *RoleService) GetRole(ns, name string) *rbacv1.Role {
	rb := svc.RoleMap.Get(ns, name)
	if rb == nil {
		panic("no such role")
	}
	return rb
}
