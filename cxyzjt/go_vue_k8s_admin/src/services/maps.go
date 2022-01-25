package services

import (
	"fmt"
	"sort"
	"sync"

	"go_vue_k8s_admin/src/models"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1beta1"
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

type V1Beta1Ingress []*v1beta1.Ingress

func (m V1Beta1Ingress) Len() int {
	return len(m)
}

func (m V1Beta1Ingress) Less(i, j int) bool {
	//根据时间排序    倒排序
	return m[i].CreationTimestamp.Time.After(m[j].CreationTimestamp.Time)
}

func (m V1Beta1Ingress) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type IngressMap struct {
	// namespace:[]*v1beta1.Ingress
	data sync.Map
}

// Add 添加Ingress
func (m *IngressMap) Add(ingress *v1beta1.Ingress) {
	if list, ok := m.data.Load(ingress.Namespace); ok {
		list = append(list.([]*v1beta1.Ingress), ingress)
		m.data.Store(ingress.Namespace, list)
	} else {
		m.data.Store(ingress.Namespace, []*v1beta1.Ingress{ingress})
	}
}

// Update 更新Ingress
func (m *IngressMap) Update(ingress *v1beta1.Ingress) error {
	if list, ok := m.data.Load(ingress.Namespace); ok {
		for i, v := range list.([]*v1beta1.Ingress) {
			if v.Name == ingress.Name {
				list.([]*v1beta1.Ingress)[i] = ingress
			}
		}
		return nil
	}
	return fmt.Errorf("ingress-%s not found", ingress.Name)
}

// Delete 删除Ingress
func (m *IngressMap) Delete(ingress *v1beta1.Ingress) {
	if list, ok := m.data.Load(ingress.Namespace); ok {
		for i, v := range list.([]*v1beta1.Ingress) {
			if v.Name == ingress.Name {
				newList := append(list.([]*v1beta1.Ingress)[:i], list.([]*v1beta1.Ingress)[i+1:]...)
				m.data.Store(ingress.Namespace, newList)
				break
			}
		}
	}
}

// Get 获取单个Ingress
func (m *IngressMap) Get(ns string, name string) *v1beta1.Ingress {
	if list, ok := m.data.Load(ns); ok {
		for _, v := range list.([]*v1beta1.Ingress) {
			if v.Name == name {
				return v
			}
		}
	}
	return nil
}

// ListAll 查询命名空间下的Ingress列表
func (m *IngressMap) ListAll(ns string) []*v1beta1.Ingress {
	if list, ok := m.data.Load(ns); ok {
		newList := list.([]*v1beta1.Ingress)
		sort.Sort(V1Beta1Ingress(newList))
		return newList
	}
	return []*v1beta1.Ingress{}
}

//func (m *IngressMap) ListAll(ns string) []*models.IngressModel {
//	if list, ok := m.data.Load(ns); ok {
//		ingressList := list.([]*v1beta1.Ingress)
//		sort.Sort(V1Beta1Ingress(ingressList))
//		result := make([]*models.IngressModel, len(ingressList))
//		for i, v := range ingressList {
//			result[i] = &models.IngressModel{
//				Name:       v.Name,
//				NameSpace:  v.Namespace,
//				CreateTime: v.CreationTimestamp.Format("2006-01-02 15:04:05"),
//				Host:       v.Spec.Rules[0].Host,
//			}
//		}
//		return result
//	}
//	return nil
//}

type CoreV1Service []*corev1.Service

func (m CoreV1Service) Len() int {
	return len(m)
}

func (m CoreV1Service) Less(i, j int) bool {
	//根据时间排序    倒排序
	return m[i].CreationTimestamp.Time.After(m[j].CreationTimestamp.Time)
}

func (m CoreV1Service) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type ServiceMap struct {
	// namespace:[]*v1.Service
	data sync.Map
}

// Add 添加Service
func (m *ServiceMap) Add(svc *corev1.Service) {
	if list, ok := m.data.Load(svc.Namespace); ok {
		list = append(list.([]*corev1.Service), svc)
		m.data.Store(svc.Namespace, list)
	} else {
		m.data.Store(svc.Namespace, []*corev1.Service{svc})
	}
}

// Update 更新Service
func (m *ServiceMap) Update(svc *corev1.Service) error {
	if list, ok := m.data.Load(svc.Namespace); ok {
		for i, v := range list.([]*corev1.Service) {
			if v.Name == svc.Name {
				list.([]*corev1.Service)[i] = svc
			}
		}
		return nil
	}
	return fmt.Errorf("service-%s not found", svc.Name)
}

// Delete 删除Service
func (m *ServiceMap) Delete(svc *corev1.Service) {
	if list, ok := m.data.Load(svc.Namespace); ok {
		for i, v := range list.([]*corev1.Service) {
			if v.Name == svc.Name {
				newList := append(list.([]*corev1.Service)[:i], list.([]*corev1.Service)[i+1:]...)
				m.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}

// ListAll 查询Service列表
func (m *ServiceMap) ListAll(ns string) []*models.ServiceModel {
	if list, ok := m.data.Load(ns); ok {
		svcList := list.([]*corev1.Service)
		sort.Sort(CoreV1Service(svcList))
		result := make([]*models.ServiceModel, len(svcList))
		for i, v := range svcList {
			result[i] = &models.ServiceModel{
				Name:       v.Name,
				NameSpace:  v.Namespace,
				CreateTime: v.CreationTimestamp.Format("2006-01-02 15:04:05"),
			}
		}
		return result
	}
	return nil
}

type CoreV1Secret []*corev1.Secret

func (m CoreV1Secret) Len() int {
	return len(m)
}
func (m CoreV1Secret) Less(i, j int) bool {
	//根据时间排序    倒排序
	return m[i].CreationTimestamp.Time.After(m[j].CreationTimestamp.Time)
}
func (m CoreV1Secret) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type SecretMap struct {
	// namespace:[]*v1.Secret
	data sync.Map
}

// Add 添加Secret
func (m *SecretMap) Add(item *corev1.Secret) {
	if list, ok := m.data.Load(item.Namespace); ok {
		list = append(list.([]*corev1.Secret), item)
		m.data.Store(item.Namespace, list)
	} else {
		m.data.Store(item.Namespace, []*corev1.Secret{item})

	}
}

// Update 更新Secret
func (m *SecretMap) Update(item *corev1.Secret) error {
	if list, ok := m.data.Load(item.Namespace); ok {
		for i, v := range list.([]*corev1.Secret) {
			if v.Name == item.Name {
				list.([]*corev1.Secret)[i] = item
			}
		}
		return nil
	}
	return fmt.Errorf("secret-%s not found", item.Name)
}

// Delete 珊瑚Secret
func (m *SecretMap) Delete(svc *corev1.Secret) {
	if list, ok := m.data.Load(svc.Namespace); ok {
		for i, v := range list.([]*corev1.Secret) {
			if v.Name == svc.Name {
				newList := append(list.([]*corev1.Secret)[:i], list.([]*corev1.Secret)[i+1:]...)
				m.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}

// Get 获取Secret
func (m *SecretMap) Get(ns string, name string) *corev1.Secret {
	if items, ok := m.data.Load(ns); ok {
		for _, v := range items.([]*corev1.Secret) {
			if v.Name == name {
				return v
			}
		}
	}
	return nil
}

// ListAll 根据命名空间获取Secret列表
func (m *SecretMap) ListAll(ns string) []*corev1.Secret {
	if list, ok := m.data.Load(ns); ok {
		secretList := list.([]*corev1.Secret)
		sort.Sort(CoreV1Secret(secretList))
		return secretList
	}
	return nil
}
