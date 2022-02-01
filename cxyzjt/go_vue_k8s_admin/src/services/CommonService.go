package services

import (
	"fmt"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

// GetImages 获取Deployment镜像字符串描述
func (svc *CommonService) GetImages(dep v1.Deployment) string {
	return svc.GetImagesByPod(dep.Spec.Template.Spec.Containers)
}

// GetImagesByPod 获取Pod镜像字符串秒速
func (svc *CommonService) GetImagesByPod(containers []corev1.Container) string {
	images := containers[0].Image
	if imgLen := len(containers); imgLen > 1 {
		images += fmt.Sprintf("+其他%d个镜像", imgLen-1)
	}
	return images
}

func (svc *CommonService) PodIsReady(pod *corev1.Pod) bool {
	if pod.Status.Phase != "Running" {
		return false
	}
	for _, condition := range pod.Status.Conditions {
		if condition.Status != "True" {
			return false
		}
	}
	for _, rg := range pod.Spec.ReadinessGates {
		for _, condition := range pod.Status.Conditions {
			if condition.Type == rg.ConditionType && condition.Status != "True" {
				return false
			}
		}
	}
	return true
}
