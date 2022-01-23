package configs

import "go_vue_k8s_admin/src/services"

type ServiceConfig struct{}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (svc *ServiceConfig) CommonService() *services.CommonService {
	return services.NewCommonService()
}

func (svc *ServiceConfig) DeploymentService() *services.DeploymentService {
	return services.NewDeploymentService()
}

func (svc *ServiceConfig) PodService() *services.PodService {
	return services.NewPodService()
}

func (*ServiceConfig) Helper() *services.Helper {
	return services.NewHelper()
}
