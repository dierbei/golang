package services

import (
	"github.com/shenyisyn/goft-gin/goft"
)

type PodService struct {
	PodMap *PodMap `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

func (svc *PodService) ListByNs(ns string) interface{} {
	list, err := svc.PodMap.ListByNS(ns)
	goft.Error(err)
	return list
}
