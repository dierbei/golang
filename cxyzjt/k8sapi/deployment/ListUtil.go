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
