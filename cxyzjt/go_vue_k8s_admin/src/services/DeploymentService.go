package services

import (
	"github.com/shenyisyn/goft-gin/goft"
	"go_vue_k8s_admin/src/core"
	"go_vue_k8s_admin/src/models"
	v1 "k8s.io/api/apps/v1"
)

type DeploymentService struct {
	DeployMap *core.DeploymentMap `inject:"-"`
	CommonSvc *CommonService      `inject:"-"`
}

func NewDeploymentService() *DeploymentService {
	return &DeploymentService{}
}

func (svc *DeploymentService) ListAll(namespace string) (result []*models.Deployment) {
	list, err := svc.DeployMap.ListByNS(namespace)
	goft.Error(err)
	for _, v := range list {
		result = append(result, &models.Deployment{Name: v.Name,
			NameSpace:  v.Namespace,
			Replicas:   [3]int32{v.Status.Replicas, v.Status.AvailableReplicas, v.Status.UnavailableReplicas},
			Images:     svc.CommonSvc.GetImages(*v),
			IsComplete: svc.getDeploymentIsComplete(v),
			Message:    svc.getDeploymentCondition(v),
		})
	}
	return
}

// getDeploymentCondition 错误消息
func (*DeploymentService) getDeploymentCondition(dep *v1.Deployment) string {
	for _, item := range dep.Status.Conditions {
		if string(item.Type) == "Available" && string(item.Status) != "True" {
			return item.Message
		}
	}
	return ""
}

// getDeploymentIsComplete 容器是否就绪
func (*DeploymentService) getDeploymentIsComplete(dep *v1.Deployment) bool {
	return dep.Status.Replicas == dep.Status.AvailableReplicas
}
