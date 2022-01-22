package core

import (
	"fmt"
	"sync"

	v1 "k8s.io/api/apps/v1"
)

type DeploymentMap struct {
	// string:[]*v1Deployment
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

func (m *DeploymentMap) ListByNS(ns string) ([]*v1.Deployment, error) {
	if list, ok := m.data.Load(ns); ok {
		return list.([]*v1.Deployment), nil
	}
	return nil, fmt.Errorf("deployment list not found")
}

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
