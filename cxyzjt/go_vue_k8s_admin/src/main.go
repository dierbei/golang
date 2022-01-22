package main

import (
	"go_vue_k8s_admin/src/configs"
	"go_vue_k8s_admin/src/controllers"
	"go_vue_k8s_admin/src/middlewares"

	"github.com/shenyisyn/goft-gin/goft"
)

func main() {
	goft.Ignite().Config(
		configs.NewK8sHandler(),    //1
		configs.NewK8sConfig(),     //2
		configs.NewK8sMaps(),       //3
		configs.NewServiceConfig(), //4
	).
		Mount("/v1",
			controllers.NewDeploymentCtl(),
			controllers.NewPodCtl(),
		).
		Attach(
			middlewares.NewCrosMiddleware(),
		).
		Launch()
}
