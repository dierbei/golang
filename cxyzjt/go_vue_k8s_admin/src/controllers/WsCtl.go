package controllers

import (
	"go_vue_k8s_admin/src/wscore"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type WsCtl struct {
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

func (ctl *WsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ws", ctl.Connect)
}
func (ctl *WsCtl) Name() string {
	return "WsCtl"
}
