package rbac

import (
	"fmt"
	"sort"
	"sync"

	rbacv1 "k8s.io/api/rbac/v1"
)

type RoleMap struct {
	// namespace:[]*v1.Role
	data sync.Map
}

// Get 获取Role
func (m *RoleMap) Get(ns string, name string) *rbacv1.Role {
	if items, ok := m.data.Load(ns); ok {
		for _, item := range items.([]*rbacv1.Role) {
			if item.Name == name {
				return item
			}
		}
	}
	return nil
}

// Add 添加Role
func (m *RoleMap) Add(item *rbacv1.Role) {
	if list, ok := m.data.Load(item.Namespace); ok {
		list = append(list.([]*rbacv1.Role), item)
		m.data.Store(item.Namespace, list)
	} else {
		m.data.Store(item.Namespace, []*rbacv1.Role{item})
	}
}

// Update 更新Role
func (m *RoleMap) Update(item *rbacv1.Role) error {
	if list, ok := m.data.Load(item.Namespace); ok {
		for i, v := range list.([]*rbacv1.Role) {
			if v.Name == item.Name {
				list.([]*rbacv1.Role)[i] = item
			}
		}
		return nil
	}
	return fmt.Errorf("role-%s not found", item.Name)
}

// Delete 删除Role
func (m *RoleMap) Delete(svc *rbacv1.Role) {
	if list, ok := m.data.Load(svc.Namespace); ok {
		for i, v := range list.([]*rbacv1.Role) {
			if v.Name == svc.Name {
				newList := append(list.([]*rbacv1.Role)[:i], list.([]*rbacv1.Role)[i+1:]...)
				m.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}

func (m *RoleMap) ListAll(ns string) []*rbacv1.Role {
	if list, ok := m.data.Load(ns); ok {
		newList := list.([]*rbacv1.Role)
		sort.Sort(V1Role(newList))
		return newList
	}
	return []*rbacv1.Role{}
}

type V1Role []*rbacv1.Role

func (m V1Role) Len() int {
	return len(m)
}

func (m V1Role) Less(i, j int) bool {
	//根据时间排序    倒排序
	return m[i].CreationTimestamp.Time.After(m[j].CreationTimestamp.Time)
}

func (m V1Role) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type V1RoleBinding []*rbacv1.RoleBinding

func (m V1RoleBinding) Len() int {
	return len(m)
}
func (m V1RoleBinding) Less(i, j int) bool {
	return m[i].CreationTimestamp.Time.After(m[j].CreationTimestamp.Time)
}
func (m V1RoleBinding) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type RoleBindingMap struct {
	// namespace:[]*v1.RoleBinding
	data sync.Map
}

func (m *RoleBindingMap) Get(ns string, name string) *rbacv1.RoleBinding {
	if items, ok := m.data.Load(ns); ok {
		for _, item := range items.([]*rbacv1.RoleBinding) {
			if item.Name == name {
				return item
			}
		}
	}
	return nil
}

func (m *RoleBindingMap) Add(item *rbacv1.RoleBinding) {
	if list, ok := m.data.Load(item.Namespace); ok {
		list = append(list.([]*rbacv1.RoleBinding), item)
		m.data.Store(item.Namespace, list)
	} else {
		m.data.Store(item.Namespace, []*rbacv1.RoleBinding{item})
	}
}

func (m *RoleBindingMap) Update(item *rbacv1.RoleBinding) error {
	if list, ok := m.data.Load(item.Namespace); ok {
		for i, v := range list.([]*rbacv1.RoleBinding) {
			if v.Name == item.Name {
				list.([]*rbacv1.RoleBinding)[i] = item
			}
		}
		return nil
	}
	return fmt.Errorf("role-%s not found", item.Name)
}

func (m *RoleBindingMap) Delete(svc *rbacv1.RoleBinding) {
	if list, ok := m.data.Load(svc.Namespace); ok {
		for i, v := range list.([]*rbacv1.RoleBinding) {
			if v.Name == svc.Name {
				newList := append(list.([]*rbacv1.RoleBinding)[:i], list.([]*rbacv1.RoleBinding)[i+1:]...)
				m.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}

func (m *RoleBindingMap) ListAll(ns string) []*rbacv1.RoleBinding {
	if list, ok := m.data.Load(ns); ok {
		newList := list.([]*rbacv1.RoleBinding)
		sort.Sort(V1RoleBinding(newList)) //  按时间倒排序
		return newList
	}
	return []*rbacv1.RoleBinding{}
}
