package router

import (
	v1 "gin-admin-back/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.Engine) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", v1.Register)
	}
}
