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

type SecretCtl struct {
	SecretMap     *services.SecretMap     `inject:"-"`
	SecretService *services.SecretService `inject:"-"`
	Client        *kubernetes.Clientset   `inject:"-"`
}

func NewSecretCtl() *SecretCtl {
	return &SecretCtl{}
}

func (*SecretCtl) Name() string {
	return "SecretCtl"
}

func (ctl *SecretCtl) RmSecret(ctx *gin.Context) goft.Json {
	ns := ctx.DefaultQuery("ns", "default")
	name := ctx.DefaultQuery("name", "")
	goft.Error(ctl.Client.CoreV1().Secrets(ns).
		Delete(ctx, name, v1.DeleteOptions{}))
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}

func (ctl *SecretCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": ctl.SecretService.ListSecret(ns), //暂时 不分页
	}
}

func (ctl *SecretCtl) PostSecret(c *gin.Context) goft.Json {
	postModel := &models.PostSecretModel{}
	err := c.ShouldBindJSON(postModel)
	goft.Error(err)
	_, err = ctl.Client.CoreV1().Secrets(postModel.NameSpace).Create(
		c,
		&corev1.Secret{
			ObjectMeta: v1.ObjectMeta{
				Name:      postModel.Name,
				Namespace: postModel.NameSpace,
			},
			Type:       corev1.SecretType(postModel.Type),
			StringData: postModel.Data,
		},
		v1.CreateOptions{},
	)
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}

func (ctl *SecretCtl) Detail(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	if ns == "" || name == "" {
		panic("error param:ns or name")
	}
	secret, err := ctl.Client.CoreV1().Secrets(ns).Get(c, name, v1.GetOptions{})
	goft.Error(err)

	return gin.H{
		"code": 20000,
		"data": &models.SecretModel{
			Name:       secret.Name,
			NameSpace:  secret.Namespace,
			Type:       string(secret.Type),
			CreateTime: secret.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Data:       secret.Data,
			ExtData:    ctl.SecretService.ParseIfTLS(string(secret.Type), secret.Data),
		},
	}
}

func (ctl *SecretCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/secrets", ctl.ListAll)
	goft.Handle("DELETE", "/secrets", ctl.RmSecret)
	goft.Handle("POST", "/secrets", ctl.PostSecret)
	goft.Handle("GET", "/secrets/:ns/:name", ctl.Detail)
}
