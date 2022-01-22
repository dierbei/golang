package services

import (
	"fmt"
	"sync"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

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
func (m *PodMap) ListByNS(ns string) ([]*corev1.Pod, error) {
	if list, ok := m.data.Load(ns); ok {
		return list.([]*corev1.Pod), nil
	}
	return nil, fmt.Errorf("pod list not found")
}
