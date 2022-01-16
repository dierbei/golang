package deployment

import (
	"k8sapi/core"

	"k8sapi/lib"
)

//显示 所有
func ListAll(namespace string) (ret []*Deployment) {
	depList, err := core.DepMap.ListByNS(namespace)
	lib.CheckError(err)
	for _, item := range depList {
		ret = append(ret, &Deployment{
			Name:     item.Name,
			Replicas: [3]int32{item.Status.Replicas, item.Status.AvailableReplicas, item.Status.UnavailableReplicas},
			Images:   GetImagesByDep(*item),
		})
	}
	return
}

//这里的函数好比 DTO . 把原生的 deployment 或pod转换为  自己的 实体对象
func ListPodsByLabel(ns string, labels map[string]string) (ret []*Pod) {
	list, err := core.PodMap.ListByLabels(ns, labels)
	lib.CheckError(err)
	for _, item := range list {
		ret=append(ret,&Pod{
			Name:item.Name,
			Images:GetImages(item.Spec.Containers),
			NodeName:item.Spec.NodeName,
			Phase:string(item.Status.Phase),// 阶段
			Message:GetPodMessage(*item),
			CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return
}
