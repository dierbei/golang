package core

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8sapi/lib"
	"log"
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
	podInformer:=fact.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(&PodHandler{})

	rsInformer:=fact.Apps().V1().ReplicaSets()
	rsInformer.Informer().AddEventHandler(&RsHandler{})
	fact.Start(wait.NeverStop)

}

func(this *DeploymentMap) GetDeployment(ns string,depname string) (*v1.Deployment,error){
	if list,ok:=this.Data.Load(ns);ok {
		for _,item:=range list.([]*v1.Deployment){
			if item.Name==depname{
				return item,nil
			}
		}
	}
	return nil,fmt.Errorf("record not found")
}

type DepHandler struct{}

func (this *DepHandler) OnAdd(obj interface{}) {
	DepMap.Add(obj.(*v1.Deployment))
}

func (this *DepHandler) OnUpdate(oldObj, newObj interface{}) {
	err := DepMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	}
}

func (this *DepHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*v1.Deployment); ok {
		DepMap.Delete(d)
	}
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

func (this *DeploymentMap) Update(deploy *v1.Deployment) error {
	if list, ok := this.Data.Load(deploy.Namespace); ok {
		for i, v := range list.([]*v1.Deployment) {
			if v.Name == deploy.Name {
				list.([]*v1.Deployment)[i] = deploy
			}
		}
		return nil
	}
	return fmt.Errorf("deployment-%s not found", deploy.Name)
}

func (this *DeploymentMap) Delete(deploy *v1.Deployment) {
	if list, ok := this.Data.Load(deploy.Namespace); ok {
		for i, v := range list.([]*v1.Deployment) {
			if v.Name == deploy.Name {
				list = append(list.([]*v1.Deployment)[:i], list.([]*v1.Deployment)[i+1:]...)
				this.Data.Store(deploy.Namespace, list)
				break
			}
		}
	}
}
