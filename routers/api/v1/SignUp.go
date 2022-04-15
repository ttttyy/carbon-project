package v1

import (
	"carbon/common"
	"carbon/lib"
	"carbon/pkg/app"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
)

func SignUp(ctx *gin.Context) {
	DB := common.GetDB()
	appG := app.Gin{C: ctx}
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	passwd := ctx.PostForm("passwd")
	if len(name) < 5 {
		appG.Response(http.StatusBadRequest, "名称太短", gin.H{
			"name": name,
		})
		return
	}
	if len(passwd) < 8 {
		appG.Response(http.StatusBadRequest, "密码强度不足", gin.H{
			"passwd": passwd,
			"name":   name,
			"email":  email,
		})
		return
	}
	if VerifyEmailFormat(email) != true {
		appG.Response(http.StatusBadRequest, "邮箱格式错误", gin.H{
			"email": email,
		})
		return
	}
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "加密错误", nil)
		return
	}
	newUser := lib.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPasswd),
		Quota:    6000,
		Money:    100000,
	}
	DB.Create(&newUser)
	appG.Response(http.StatusOK, "注册成功", gin.H{
		"msg":    "SignUp success",
		"name":   name,
		"email":  email,
		"passwd": hashedPasswd,
	})
}
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
