package controllers

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type PodLogsCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
}

func NewPodLogsCtl() *PodLogsCtl {
	return &PodLogsCtl{}
}

func (ctl *PodLogsCtl) GetLogs(c *gin.Context) (v goft.Void) {
	ns := c.DefaultQuery("ns", "default")
	podname := c.DefaultQuery("podname", "")
	cname := c.DefaultQuery("cname", "")
	var tailLine int64 = 100
	opt := &v1.PodLogOptions{Follow: true, Container: cname, TailLines: &tailLine}
	cc, _ := context.WithTimeout(c, time.Minute*30) //设置半小时超时时间。否则会造成内存泄露

	req := ctl.Client.CoreV1().Pods(ns).GetLogs(podname, opt)
	reader, err := req.Stream(cc)
	defer reader.Close()
	goft.Error(err)
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		if n > 0 {
			c.Writer.Write([]byte(string(buf[0:n])))
			c.Writer.(http.Flusher).Flush()
		}
	}
	return
}

func (*PodLogsCtl) Name() string {
	return "PodLogsCtl"
}

func (ctl *PodLogsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods/logs", ctl.GetLogs)
}
