package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"mypro/src/controllers"
	"net/http"
)
func crosJwks() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}else{
			c.Next()
		}
	}
}

func main() {
	goft.Ignite(crosJwks()).Mount("",
		controllers.NewJwksCtl(),
	).Launch()
}