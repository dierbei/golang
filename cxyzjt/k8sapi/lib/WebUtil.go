package lib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckError(err error){
	if err!=nil{
		panic(err.Error())
	}
}

func Success(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}