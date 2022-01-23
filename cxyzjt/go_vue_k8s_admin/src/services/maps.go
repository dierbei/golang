package services

import (
	"fmt"
	"go_vue_k8s_admin/src/models"
	"sort"
	"sync"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type MapItems []*MapItem

type MapItem struct {
	key   string
	value interface{}
}

func convertToMapItems(m sync.Map) MapItems {
	items := make(MapItems, 0)
	m.Range(func(key, value interface{}) bool {
		items = append(items, &MapItem{key: key.(string), value: value})
		return true
	})
	return items
}

func (m MapItems) Len() int {
	return len(m)
}
func (m MapItems) Less(i, j int) bool {
	return m[i].key < m[j].key
}
func (m MapItems) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type DeploymentMap struct {
	// namespace:[]*v1Deployment
	data sync.Map
}

// Add 添加Deployment
func (m *DeploymentMap) Add(deploy *v1.Deployment) {
	if list, ok := m.data.Load(deploy.Namespace); ok {
		list = append(list.([]*v1.Deployment), deploy)
		m.data.Store(deploy.Namespace, list)
	} else {
		m.data.Store(deploy.Namespace, []*v1.Deployment{deploy})
	}
}

// Update 更新Deployment
func (m *DeploymentMap) Update(deploy *v1.Deployment) error {
	if list, ok := m.data.Load(deploy.Namespace); ok {
		for i, v := range list.([]*v1.Deployment) {
			if v.Name == deploy.Name {
				list.([]*v1.Deployment)[i] = deploy
			}
			return nil
		}
	}

	return fmt.Errorf("deployment-%s not found", deploy.Name)
}

// Delete 删除Deployment
func (m *DeploymentMap) Delete(deploy *v1.Deployment) {
	if list, ok := m.data.Load(deploy.Namespace); ok {
		for i, v := range list.([]*v1.Deployment) {
			if v.Name == deploy.Name {
				list = append(list.([]*v1.Deployment)[:i], list.([]*v1.Deployment)[i+1:]...)
				m.data.Store(deploy.Namespace, list)
				break
			}
		}
	}
}

// ListByNS 根据命名空间查询Deployment列表
func (m *DeploymentMap) ListByNS(ns string) ([]*v1.Deployment, error) {
	if list, ok := m.data.Load(ns); ok {
		return list.([]*v1.Deployment), nil
	}
	return nil, fmt.Errorf("deployment list not found")
}

// GetDeployment 根据命名空间和deployment查询Deployment
func (m *DeploymentMap) GetDeployment(ns string, deployName string) (*v1.Deployment, error) {
	if list, ok := m.data.Load(ns); ok {
		for _, v := range list.([]*v1.Deployment) {
			if v.Name == deployName {
				return v, nil
			}
		}
	}
	return nil, fmt.Errorf("deployment not found")
}

type CoreV1Pods []*corev1.Pod

func (c CoreV1Pods) Len() int {
	return len(c)
}

func (c CoreV1Pods) Less(i, j int) bool {
	//根据时间排序    正排序
	return c[i].CreationTimestamp.Time.Before(c[j].CreationTimestamp.Time)
}

func (c CoreV1Pods) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type PodMap struct {
	// namespace:[]*v1.Pod
	data sync.Map
}

// Add 添加Pod
func (m *PodMap) Add(pod *corev1.Pod) {
	if list, ok := m.data.Load(pod.Namespace); ok {
		list = append(list.([]*corev1.Pod), pod)
		m.data.Store(pod.Namespace, list)
	} else {
		m.data.Store(pod.Namespace, []*corev1.Pod{pod})
	}
}

// Update 更新Pod
func (m *PodMap) Update(pod *corev1.Pod) error {
	if list, ok := m.data.Load(pod.Namespace); ok {
		for i, v := range list.([]*corev1.Pod) {
			if v.Name == pod.Name {
				list.([]*corev1.Pod)[i] = pod
			}
			return nil
		}
	}

	return fmt.Errorf("pod-%s not found", pod.Name)
}

// Delete 删除Pod
func (m *PodMap) Delete(pod *corev1.Pod) {
	if list, ok := m.data.Load(pod.Namespace); ok {
		for i, v := range list.([]*corev1.Pod) {
			if v.Name == pod.Name {
				list = append(list.([]*corev1.Pod)[:i], list.([]*corev1.Pod)[i+1:]...)
				m.data.Store(pod.Namespace, list)
				break
			}
		}
	}
}

// ListByNS 根据命名空间查询Deployment列表
func (m *PodMap) ListByNS(ns string) []*corev1.Pod {
	if list, ok := m.data.Load(ns); ok {
		result := list.([]*corev1.Pod)
		sort.Sort(CoreV1Pods(result))
		return result
	}
	return nil
}

type NameSpaceMap struct {
	// name:[]*corev1.Namespace
	data sync.Map
}

// Add 添加NameSpace
func (m *NameSpaceMap) Add(ns *corev1.Namespace) {
	m.data.Store(ns.Name, ns)
}

// Update 更新NameSpace
func (m *NameSpaceMap) Update(ns *corev1.Namespace) {
	m.data.Store(ns.Name, ns)
}

// Delete 删除NameSpace
func (m *NameSpaceMap) Delete(ns *corev1.Namespace) {
	m.data.Delete(ns.Name)
}

// Get 获取NameSpace
func (m *NameSpaceMap) Get(name string) *corev1.Namespace {
	if item, ok := m.data.Load(name); ok {
		return item.(*corev1.Namespace)
	}
	return nil
}

// ListAll 获取所有NameSpace
func (m *NameSpaceMap) ListAll() []*models.NamespaceModel {
	items := convertToMapItems(m.data)
	sort.Sort(items)
	nsList := make([]*models.NamespaceModel, len(items))
	for i, v := range items {
		nsList[i] = &models.NamespaceModel{Name: v.key}
	}
	return nsList
}

type EventMap struct {
	// namespace_king_name:*v1.Event
	data sync.Map
}

// GetMessage 获取
func (m *EventMap) GetMessage(ns string, king string, name string) string {
	key := fmt.Sprintf("%s_%s_%s", ns, king, name)
	if v, ok := m.data.Load(key); ok {
		return v.(*corev1.Event).Message
	}
	return ""
}
