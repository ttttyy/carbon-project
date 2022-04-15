package v1

import (
	"carbon/common"
	"carbon/lib"
	"carbon/pkg/app"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignIn(ctx *gin.Context) {
	DB := common.GetDB()
	appG := app.Gin{C: ctx}
	name := ctx.PostForm("name")
	passwd := ctx.PostForm("passwd")
	var user lib.User
	DB.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		appG.Response(http.StatusUnprocessableEntity, "用户不存在", gin.H{
			"h1": user.Password,
			"h3": user.Name,
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwd)); err != nil {
		appG.Response(http.StatusBadRequest, "密码错误", gin.H{
			"h1": user.Password,
			"h3": user.Name,
		})
		return
	}
	token, err := common.GenerateToken(user)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "系统异常", gin.H{
			"code": 500,
		})
	}
	ctx.JSON(200, gin.H{
		"msg":    "SignIn success",
		"name":   name,
		"passwd": passwd,
		"token":  token,
	})
}
