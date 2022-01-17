package deployment

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sapi/lib"

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

//根据dep 获取 当前rs的标签
func GetRsLabelByDeployment(dep *v1.Deployment) string {
	selector, _ := metav1.LabelSelectorAsSelector(dep.Spec.Selector)
	listOpt := metav1.ListOptions{
		LabelSelector: selector.String(),
	}
	rs, _ := lib.K8sClient.AppsV1().ReplicaSets(dep.Namespace).
		List(context.Background(), listOpt)
	for _, item := range rs.Items {
		if IsCurrentRsByDep(dep, item) {
			s, err := metav1.LabelSelectorAsSelector(item.Spec.Selector)
			if err != nil {
				return ""
			}
			return s.String()
		}
	}
	return ""
}

func IsCurrentRsByDep(dep *v1.Deployment, set v1.ReplicaSet) bool {
	if set.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] != dep.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] {
		return false
	}
	for _, ref := range set.OwnerReferences {
		if ref.Kind == "Deployment" && ref.Name == dep.Name {
			return true
		}
	}
	return false
}

//根据dep 获取 当前rs的标签  --- 给listwatch使用
func GetRsLableByDeployment_ListWatch(dep *v1.Deployment,rslist []*v1.ReplicaSet) (map[string]string,error){
	for _,item:=range rslist{
		if IsCurrentRsByDep(dep,*item){
			s,err:= metav1.LabelSelectorAsMap(item.Spec.Selector)
			if err!=nil{
				return nil,err
			}
			return s,nil
		}
	}
	return nil,nil
}