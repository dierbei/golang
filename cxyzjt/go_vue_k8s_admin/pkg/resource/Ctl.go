package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"strings"
)

type ResourcesCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
}

func NewResourcesCtl() *ResourcesCtl {
	return &ResourcesCtl{}
}

func (ctl *ResourcesCtl) GetGroupVersion(str string) (group, version string) {
	list := strings.Split(str, "/")
	if len(list) == 1 {
		return "core", list[0]
	} else if len(list) == 2 {
		return list[0], list[1]
	}
	panic("error GroupVersion" + str)
}

func (ctl *ResourcesCtl) ListResources(c *gin.Context) goft.Json {
	_, apiResourceList, err := ctl.Client.ServerGroupsAndResources()
	goft.Error(err)
	gRes := make([]*GroupResources, 0)
	for _, resource := range apiResourceList {
		group, version := ctl.GetGroupVersion(resource.GroupVersion)
		gr := &GroupResources{Group: group, Version: version, Resources: make([]*Resources, 0)}
		for _, rr := range resource.APIResources {
			res := &Resources{Name: rr.Name, Verbs: rr.Verbs}
			gr.Resources = append(gr.Resources, res)
		}
		gRes = append(gRes, gr)
	}
	return gin.H{
		"code": 20000,
		"data": gRes,
	}
}


func (*ResourcesCtl) Name() string {
	return "Resources"
}

func (ctl *ResourcesCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/resources", ctl.ListResources)
}
