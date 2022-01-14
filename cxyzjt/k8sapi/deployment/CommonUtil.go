package deployment

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
)

func GetImages(containers []core.Container) string {
	image := containers[0].Image
	if imgLen := len(containers); imgLen > 1 {
		image += fmt.Sprintf("和其它%d个镜像", imgLen-1)
	}
	return image
}

func GetImagesByDep(deployment v1.Deployment) string {
	return GetImages(deployment.Spec.Template.Spec.Containers)
}

func GetLabels(m map[string]string) string {
	labels := ""
	for k, v := range m {
		if labels != "" {
			labels += ","
		}
		labels += fmt.Sprintf("%s=%s", k, v)
	}
	return labels
}