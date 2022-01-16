package deployment

import (
	"context"

	"k8sapi/lib"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Detail(namespace string, depName string) *Deployment {
	listopt := metav1.GetOptions{}
	deployment, err := lib.K8sClient.AppsV1().Deployments(namespace).Get(context.Background(), depName, listopt)
	lib.CheckError(err)
	return &Deployment{
		Name:       deployment.Name,
		NameSpace:  deployment.Namespace,
		Images:     GetImagesByDep(*deployment),
		Pods:       GetPodByDep(namespace, deployment),
		Replicas:   [3]int32{deployment.Status.Replicas, deployment.Status.AvailableReplicas, deployment.Status.UnavailableReplicas},
		CreateTime: deployment.CreationTimestamp.Format("2006-01-02 15:04:05"),
	}
}

func GetPodByDep(namespace string, dep *v1.Deployment) []Pod {
	getopt := metav1.ListOptions{LabelSelector: GetRsLabelByDeployment(dep)}
	podList, err := lib.K8sClient.CoreV1().Pods(namespace).List(context.Background(), getopt)
	lib.CheckError(err)
	podSlice := make([]Pod, len(podList.Items))
	for index, item := range podList.Items {
		pod := Pod{
			Name:       item.Name,
			Images:     GetImages(item.Spec.Containers),
			NodeName:   item.Spec.NodeName,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
		podSlice[index] = pod
	}
	return podSlice
}
