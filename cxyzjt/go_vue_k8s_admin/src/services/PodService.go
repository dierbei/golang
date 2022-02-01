package services

import (
	"github.com/gin-gonic/gin"
	"go_vue_k8s_admin/src/models"
)

type PodService struct {
	PodMap    *PodMap        `inject:"-"`
	CommonSvc *CommonService `inject:"-"`
	EventMap  *EventMap      `inject:"-"`
	Helper    Helper         `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

func (svc *PodService) ListByNs(ns string) interface{} {
	podList := svc.PodMap.ListByNS(ns)
	ret := make([]*models.Pod, 0)
	for _, pod := range podList {
		ret = append(ret, &models.Pod{
			Name:       pod.Name,
			NameSpace:  pod.Namespace,
			Images:     svc.CommonSvc.GetImagesByPod(pod.Spec.Containers),
			NodeName:   pod.Spec.NodeName,
			Phase:      string(pod.Status.Phase),      // 阶段
			IsReady:    svc.CommonSvc.PodIsReady(pod), //是否就绪
			IP:         []string{pod.Status.PodIP, pod.Status.HostIP},
			Message:    svc.EventMap.GetMessage(pod.Namespace, "Pod", pod.Name),
			CreateTime: pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return ret
}

func (svc *PodService) PagePods(ns string, page, size int) *ItemsPage {
	pods := svc.ListByNs(ns).([]*models.Pod)
	readyCount := 0 //就绪的pod数量
	allCount := 0   //总数量
	ipods := make([]interface{}, len(pods))
	for i, pod := range pods {
		allCount++
		ipods[i] = pod
		if pod.IsReady {
			readyCount++
		}
	}
	return svc.Helper.PageResource(
		page,
		size,
		ipods).SetExt(gin.H{"ReadyNum": readyCount, "AllNum": allCount})
}

func (svc *PodService) GetPodContainer(ns, podname string) []*models.ContainerModel {
	ret := make([]*models.ContainerModel, 0)
	pod := svc.PodMap.Get(ns, podname)
	if pod != nil {
		for _, c := range pod.Spec.Containers {
			ret = append(ret, &models.ContainerModel{
				Name: c.Name,
			})
		}
	}
	return ret
}
