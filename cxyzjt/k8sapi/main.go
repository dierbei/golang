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
		  c.HTML(http.StatusOK, "list.html",
		  		  lib.DataBuilder().
		  			SetTitle("deployment列表").
		  			SetData("DepList",deployment.ListAll("myweb")))
	  })

      r.Run(":8080")
}
