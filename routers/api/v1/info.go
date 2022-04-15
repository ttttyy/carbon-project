package v1

import (
	"carbon/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	appG := app.Gin{C: ctx}
	appG.Response(http.StatusOK, "获取成功", gin.H{
		"user": user,
	})
}
