package main

import (
	"github.com/gin-gonic/gin"
	"k8sapi/deployment"
	"k8sapi/lib"
	"net/http"
)

func main() {
      r:=gin.New()
	r.Static("/static", "./static")
	  r.LoadHTMLGlob("html/**/*")

      r.GET("/deployments", func(c *gin.Context) {
		  c.HTML(http.StatusOK, "deployment_list.html",
		  		  lib.DataBuilder().
		  			SetTitle("deployment列表").
		  			SetData("DepList",deployment.ListAll("myweb")))
	  })
	r.GET("/deployments/:name", func(c *gin.Context) {
		c.HTML(http.StatusOK, "deployment_detail.html",
			lib.DataBuilder().
				SetTitle("deployment详情").
				SetData("DepDetail",deployment.Detail("myweb", c.Param("name"))))
	})

      r.Run(":8080")
}
