package rbac

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
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

// ListAll 获取Role列表
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

// Get 获取RoleBinding
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

// Add 添加RoleBinding
func (m *RoleBindingMap) Add(item *rbacv1.RoleBinding) {
	if list, ok := m.data.Load(item.Namespace); ok {
		list = append(list.([]*rbacv1.RoleBinding), item)
		m.data.Store(item.Namespace, list)
	} else {
		m.data.Store(item.Namespace, []*rbacv1.RoleBinding{item})
	}
}

// Update 更新RoleBinding
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

// Delete 删除RoleBinding
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

// ListAll 查询RoleBinding列表
func (m *RoleBindingMap) ListAll(ns string) []*rbacv1.RoleBinding {
	if list, ok := m.data.Load(ns); ok {
		newList := list.([]*rbacv1.RoleBinding)
		sort.Sort(V1RoleBinding(newList)) //  按时间倒排序
		return newList
	}
	return []*rbacv1.RoleBinding{}
}

type CoreV1Sa []*corev1.ServiceAccount

func (m CoreV1Sa) Len() int {
	return len(m)
}

func (m CoreV1Sa) Less(i, j int) bool {
	return m[i].CreationTimestamp.Time.After(m[j].CreationTimestamp.Time)
}

func (m CoreV1Sa) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type ServiceAccountMap struct {
	// namesapce:[]*corev1.ServiceAccount
	data sync.Map
}

// Get 获取ServiceAccount
func (m *ServiceAccountMap) Get(ns string, name string) *corev1.ServiceAccount {
	if items, ok := m.data.Load(ns); ok {
		for _, item := range items.([]*corev1.ServiceAccount) {
			if item.Name == name {
				return item
			}
		}
	}
	return nil
}

// Add 添加ServiceAccount
func (m *ServiceAccountMap) Add(item *corev1.ServiceAccount) {
	if list, ok := m.data.Load(item.Namespace); ok {
		list = append(list.([]*corev1.ServiceAccount), item)
		m.data.Store(item.Namespace, list)
	} else {
		m.data.Store(item.Namespace, []*corev1.ServiceAccount{item})
	}
}

// Update 更新ServiceAccount
func (m *ServiceAccountMap) Update(item *corev1.ServiceAccount) error {
	if list, ok := m.data.Load(item.Namespace); ok {
		for i, v := range list.([]*corev1.ServiceAccount) {
			if v.Name == item.Name {
				list.([]*corev1.ServiceAccount)[i] = item
			}
		}
		return nil
	}
	return fmt.Errorf("ServiceAccount-%s not found", item.Name)
}

// Delete 删除ServiceAccount
func (m *ServiceAccountMap) Delete(sa *corev1.ServiceAccount) {
	if list, ok := m.data.Load(sa.Namespace); ok {
		for i, v := range list.([]*corev1.ServiceAccount) {
			if v.Name == sa.Name {
				newList := append(list.([]*corev1.ServiceAccount)[:i], list.([]*corev1.ServiceAccount)[i+1:]...)
				m.data.Store(sa.Namespace, newList)
				break
			}
		}
	}
}

// ListAll 获取ServiceAccount列表
func (m *ServiceAccountMap) ListAll(ns string) []*corev1.ServiceAccount {
	if list, ok := m.data.Load(ns); ok {
		newList := list.([]*corev1.ServiceAccount)
		sort.Sort(CoreV1Sa(newList))
		return newList
	}
	return []*corev1.ServiceAccount{}
}

type V1ClusterRole []*rbacv1.ClusterRole

func (m V1ClusterRole) Len() int {
	return len(m)
}

func (m V1ClusterRole) Less(i, j int) bool {
	return m[i].CreationTimestamp.Time.After(m[j].CreationTimestamp.Time)
}

func (m V1ClusterRole) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type ClusterRoleMap struct {
	// name:*rbacv1.ClusterRole
	data sync.Map
}

func (m *ClusterRoleMap) Get(name string) *rbacv1.ClusterRole {
	if item, ok := m.data.Load(name); ok {
		return item.(*rbacv1.ClusterRole)
	}
	return nil
}

func (m *ClusterRoleMap) Add(item *rbacv1.ClusterRole) {
	m.data.Store(item.Name, item)
}

func (m *ClusterRoleMap) Update(item *rbacv1.ClusterRole) error {
	m.data.Store(item.Name, item)
	return nil
}

func (m *ClusterRoleMap) Delete(svc *rbacv1.ClusterRole) {
	m.data.Delete(svc.Name)

}

func (m *ClusterRoleMap) ListAll() []*rbacv1.ClusterRole {
	var list []*rbacv1.ClusterRole
	m.data.Range(func(key, value interface{}) bool {
		list = append(list, value.(*rbacv1.ClusterRole))
		return true
	})
	sort.Sort(V1ClusterRole(list))
	return list
}

type V1ClusterRoleBinding []*rbacv1.ClusterRoleBinding

func (m V1ClusterRoleBinding) Len() int {
	return len(m)
}

func (m V1ClusterRoleBinding) Less(i, j int) bool {
	//根据时间排序    倒排序
	return m[i].CreationTimestamp.Time.After(m[j].CreationTimestamp.Time)
}

func (m V1ClusterRoleBinding) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type ClusterRoleBindingMap struct {
	// name:*v1.ClusterRoleBinding
	data sync.Map //
}

func (m *ClusterRoleBindingMap) Get(name string) *rbacv1.ClusterRoleBinding {
	if v, ok := m.data.Load(name); ok {
		return v.(*rbacv1.ClusterRoleBinding)
	}
	return nil
}

func (m *ClusterRoleBindingMap) Add(item *rbacv1.ClusterRoleBinding) {
	m.data.Store(item.Name, item)
}

func (m *ClusterRoleBindingMap) Update(item *rbacv1.ClusterRoleBinding) error {
	m.data.Store(item.Name, item)
	return nil
}

func (m *ClusterRoleBindingMap) Delete(svc *rbacv1.ClusterRoleBinding) {
	m.data.Delete(svc.Name)
}

func (m *ClusterRoleBindingMap) ListAll() []*rbacv1.ClusterRoleBinding {
	var list []*rbacv1.ClusterRoleBinding
	m.data.Range(func(key, value interface{}) bool {
		list = append(list, value.(*rbacv1.ClusterRoleBinding))
		return true
	})
	sort.Sort(V1ClusterRoleBinding(list))
	return list
}
