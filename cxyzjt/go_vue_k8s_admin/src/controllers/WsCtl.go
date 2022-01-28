package controllers

import (
	"go_vue_k8s_admin/src/helpers"
	"go_vue_k8s_admin/src/wscore"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type WsCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
	Config *rest.Config          `inject:"-"`
}

func NewWsCtl() *WsCtl {
	return &WsCtl{}
}

func (ctl *WsCtl) Connect(ctx *gin.Context) string {
	client, err := wscore.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return err.Error()
	} else {
		wscore.ClientMap.Store(client)
		return "success"
	}
}

func (ctl *WsCtl) PodConnect(c *gin.Context) (v goft.Void) {
	ns := c.Query("ns")
	pod := c.Query("pod")
	container := c.Query("c")
	wsClient, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	shellClient := wscore.NewWsShellClient(wsClient)
	err = helpers.HandleCommand(ns, pod, container, ctl.Client, ctl.Config, []string{"sh"}).
		Stream(remotecommand.StreamOptions{
			Stdin:  shellClient,
			Stdout: shellClient,
			Stderr: shellClient,
			Tty:    true,
		})
	if err != nil {
		log.Println(err)
	}
	return
}

func (ctl *WsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ws", ctl.Connect)
	goft.Handle("GET", "/podws", ctl.PodConnect)
}
func (ctl *WsCtl) Name() string {
	return "WsCtl"
}
