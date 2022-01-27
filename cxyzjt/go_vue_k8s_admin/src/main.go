package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_vue_k8s_admin/src/configs"
	"go_vue_k8s_admin/src/controllers"

	"github.com/shenyisyn/goft-gin/goft"
)

func main() {
	goft.Ignite(cross()).Config(
		configs.NewK8sHandler(),    //1
		configs.NewK8sConfig(),     //2
		configs.NewK8sMaps(),       //3
		configs.NewServiceConfig(), //4
	).
		Mount("",
			controllers.NewDeploymentCtl(),
			controllers.NewPodCtl(),
			controllers.NewUserCtl(),
			controllers.NewWsCtl(),
			controllers.NewNamespaceCtl(),
			controllers.NewIngressCtl(),
			controllers.NewSvcCtl(),
			controllers.NewSecretCtl(),
			controllers.NewConfigMapCtl(),
		).
		//Attach(
		//	middlewares.NewCrosMiddleware(),
		//).
		Launch()
}

func cross() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

	}
}
