package routers

import (
	v1 "carbon/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	apiV1 := r.Group("api/v1")
	{
		apiV1.POST("/signup", v1.SignUp)
		apiV1.POST("/signin", v1.SignIn)
		apiV1.POST("/transaction", v1.Transaction)
		apiV1.POST("/future", v1.Future)
	}
	return r
}
