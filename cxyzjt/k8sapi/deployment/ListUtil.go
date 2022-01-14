package deployment

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8sapi/lib"
)

type Deployment struct {
	Name     string
	Replicas [3]int32
}

//显示 所有
func ListAll(namespace string) (ret []*Deployment) {
	ctx := context.Background()
	listopt := metav1.ListOptions{}
	depList, err := lib.K8sClient.AppsV1().Deployments(namespace).List(ctx, listopt)
	lib.CheckError(err)
	for _, item := range depList.Items { //遍历所有deployment
		ret = append(ret, &Deployment{
			Name: item.Name,
			Replicas: [3]int32{item.Status.Replicas, item.Status.AvailableReplicas, item.Status.UnavailableReplicas},
		})
	}
	return
}
