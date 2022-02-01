package services

import (
	"go_vue_k8s_admin/src/models"

	"k8s.io/client-go/kubernetes"
)

type ConfigMapService struct {
	Client    *kubernetes.Clientset `inject:"-"`
	ConfigMap *ConfigMap            `inject:"-"`
}

func NewConfigMapService() *ConfigMapService {
	return &ConfigMapService{}
}

func (svc *ConfigMapService) ListConfigMap(ns string) []*models.ConfigMapModel {
	list := svc.ConfigMap.ListAll(ns)
	ret := make([]*models.ConfigMapModel, len(list))
	for i, item := range list {
		ret[i] = &models.ConfigMapModel{
			Name:       item.Name,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:  item.Namespace,
		}
	}
	return ret
}
