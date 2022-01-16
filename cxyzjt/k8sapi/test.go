package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sapi/lib"
	"log"
)

func main() {
	dep,_:=lib.K8sClient.AppsV1().Deployments("default").
		Get(context.Background(),"ngx1",metav1.GetOptions{})
	selector,err:=metav1.LabelSelectorAsSelector(dep.Spec.Selector)
	if err!=nil{
		log.Fatal(err)
	}
	listOpt:=metav1.ListOptions{
		LabelSelector:selector.String(),
	}
	rs,_:=lib.K8sClient.AppsV1().ReplicaSets("default").
		List(context.Background(),listOpt)
	fmt.Println(dep.ObjectMeta.Annotations["deployment.kubernetes.io/revision"])
	for _,item:=range rs.Items{
		fmt.Println(item.Name)
		fmt.Println(IsCurrentRsByDep(dep,item))
		//fmt.Println(item.OwnerReferences)
		//fmt.Println(item.ObjectMeta.Annotations["deployment.kubernetes.io/revision"])
	}
}
func IsCurrentRsByDep(dep *v1.Deployment,set v1.ReplicaSet) bool{
	if set.ObjectMeta.Annotations["deployment.kubernetes.io/revision"]!=dep.ObjectMeta.Annotations["deployment.kubernetes.io/revision"]{
		return false
	}
	for _,ref:=range set.OwnerReferences{
		if ref.Kind=="Deployment" && ref.Name==dep.Name{
			return true
		}
	}
	return false
}


