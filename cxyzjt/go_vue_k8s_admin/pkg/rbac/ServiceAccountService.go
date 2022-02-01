package rbac

import corev1 "k8s.io/api/core/v1"

type ServiceAccountService struct {
	SaMap *ServiceAccountMap `inject:"-"`
}

func NewSaService() *ServiceAccountService {
	return &ServiceAccountService{}
}

func (svc *ServiceAccountService) ListSa(ns string) []*corev1.ServiceAccount {
	return svc.SaMap.ListAll(ns)
}
