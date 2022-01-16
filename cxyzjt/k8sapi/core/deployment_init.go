package core

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8sapi/lib"
	"sync"

	v1 "k8s.io/api/apps/v1"
)

var (
	DepMap *DeploymentMap
)

func init() {
	DepMap = &DeploymentMap{}
}

func InitDeployment() {
	fact := informers.NewSharedInformerFactory(lib.K8sClient, 0)
	depInformer := fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(&DepHandler{})
	fact.Start(wait.NeverStop)

}

type DepHandler struct{}

func (this *DepHandler) OnAdd(obj interface{}) {
	DepMap.Add(obj.(*v1.Deployment))
}

func (this *DepHandler) OnUpdate(oldObj, newObj interface{}) {
	if dep, ok := newObj.(*v1.Deployment); ok {
		fmt.Println(dep.Name)
	}
}

func (this *DepHandler) OnDelete(obj interface{}) {

}

type DeploymentMap struct {
	Data sync.Map
}

func (this *DeploymentMap) Add(deploy *v1.Deployment) {
	if list, ok := this.Data.Load(deploy.Namespace); ok {
		list = append(list.([]*v1.Deployment), deploy)
		this.Data.Store(deploy.Namespace, list)
	} else {
		this.Data.Store(deploy.Namespace, []*v1.Deployment{deploy})
	}
}

func (this *DeploymentMap) ListByNS(namespace string) ([]*v1.Deployment, error) {
	if list, ok := this.Data.Load(namespace); ok {
		return list.([]*v1.Deployment), nil
	}
	return nil, fmt.Errorf("record not found")
}
