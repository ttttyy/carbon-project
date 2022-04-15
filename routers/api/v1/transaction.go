package v1

import (
	"carbon/pkg/app"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Transaction(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	amount := ctx.PostForm("amount")
	option := ctx.PostForm("option")
	fmt.Println("ssss")
	appG.Response(http.StatusOK, "委托成功", gin.H{
		"time":   time.Now(),
		"name":   name,
		"price":  price,
		"amount": amount,
		"option": option,
	})

}
