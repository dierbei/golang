package deployment

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sapi/lib"
)

func Detail(namespace string, depName string) *Deployment {
	listopt := metav1.GetOptions{}
	deployment, err := lib.K8sClient.AppsV1().Deployments(namespace).Get(context.Background(), depName, listopt)
	lib.CheckError(err)
	fmt.Println(deployment.CreationTimestamp.Format("2006-01-02 15:04:05"))
	return &Deployment{
		Name:      deployment.Name,
		NameSpace: deployment.Namespace,
		Images:    GetImages(*deployment),
		Pods:      GetPodByDep(namespace, deployment),
	}
}

func GetPodByDep(namespace string, dep *v1.Deployment) []Pod {
	getopt := metav1.ListOptions{LabelSelector: GetLabels(dep.Spec.Selector.MatchLabels)}
	podList, err := lib.K8sClient.CoreV1().Pods(namespace).List(context.Background(), getopt)
	lib.CheckError(err)
	podSlice := make([]Pod, len(podList.Items))
	for index, item := range podList.Items {
		pod := Pod{Name: item.Name}
		podSlice[index] = pod
	}
	return podSlice
}
