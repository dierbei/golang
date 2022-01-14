package deployment

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
)

func GetImages(deployment v1.Deployment) string {
	image := deployment.Spec.Template.Spec.Containers[0].Image
	if imgLen := len(deployment.Spec.Template.Spec.Containers); imgLen > 1 {
		image += fmt.Sprintf("和其它%d个镜像", imgLen-1)
	}
	return image
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