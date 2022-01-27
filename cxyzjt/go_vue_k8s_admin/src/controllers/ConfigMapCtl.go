package controllers

import (
	"go_vue_k8s_admin/src/models"
	"go_vue_k8s_admin/src/services"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapCtl struct {
	ConfigMap     *services.ConfigMap        `inject:"-"`
	ConfigService *services.ConfigMapService `inject:"-"`
	Client        *kubernetes.Clientset      `inject:"-"`
}

func NewConfigMapCtl() *ConfigMapCtl {
	return &ConfigMapCtl{}
}
func (*ConfigMapCtl) Name() string {
	return "ConfigMapCtl"
}

func (ctl *ConfigMapCtl) PostConfigmap(c *gin.Context) goft.Json {
	postModel := &models.PostConfigMapModel{}
	err := c.ShouldBindJSON(postModel)
	goft.Error(err)

	cm := &corev1.ConfigMap{
		ObjectMeta: v1.ObjectMeta{
			Name:      postModel.Name,
			Namespace: postModel.NameSpace,
		},
		Data: postModel.Data,
	}
	if postModel.IsUpdate {
		_, err = ctl.Client.CoreV1().ConfigMaps(postModel.NameSpace).Update(c, cm, v1.UpdateOptions{})
	} else {
		_, err = ctl.Client.CoreV1().ConfigMaps(postModel.NameSpace).Create(c, cm, v1.CreateOptions{})
	}

	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}

func (ctl *ConfigMapCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": ctl.ConfigService.ListConfigMap(ns),
	}
}

func (ctl *ConfigMapCtl) RmCm(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	name := c.DefaultQuery("name", "")
	goft.Error(ctl.Client.CoreV1().ConfigMaps(ns).
		Delete(c, name, v1.DeleteOptions{}))
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}

func (ctl *ConfigMapCtl) Detail(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	if ns == "" || name == "" {
		panic("error param:ns or name")
	}
	cm, err := ctl.Client.CoreV1().ConfigMaps(ns).Get(c, name, v1.GetOptions{})
	goft.Error(err)

	return gin.H{
		"code": 20000,
		"data": &models.ConfigMapModel{
			Name:       cm.Name,
			NameSpace:  cm.Namespace,
			CreateTime: cm.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Data:       cm.Data,
		},
	}
}

func (ctl *ConfigMapCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/configmaps", ctl.ListAll)
	goft.Handle("GET", "/configmaps/:ns/:name", ctl.Detail)
	goft.Handle("DELETE", "/configmaps", ctl.RmCm)
	goft.Handle("POST", "/configmaps", ctl.PostConfigmap)
}
