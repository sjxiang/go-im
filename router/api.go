package router

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/go-im/service"
)


func registerApiRoutes(router *gin.Engine) {
	
	// 用户登录
	router.POST("/login", service.Login)

	
}
