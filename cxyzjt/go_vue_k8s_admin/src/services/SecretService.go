package services

import (
	"go_vue_k8s_admin/src/models"

	"k8s.io/client-go/kubernetes"
)

type SecretService struct {
	Client    *kubernetes.Clientset `inject:"-"`
	SecretMap *SecretMap            `inject:"-"`
}

func NewSecretService() *SecretService {
	return &SecretService{}
}

func (svc *SecretService) ListSecret(ns string) []*models.SecretModel {
	list := svc.SecretMap.ListAll(ns)
	ret := make([]*models.SecretModel, len(list))
	for i, item := range list {
		ret[i] = &models.SecretModel{
			Name:       item.Name,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:  item.Namespace,
			Type:       models.SECRET_TYPE[string(item.Type)], // 类型的翻译
		}
	}
	return ret
}
