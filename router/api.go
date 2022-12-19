package router

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/go-im/middleware"
	"github.com/sjxiang/go-im/service"
)


func registerApiRoutes(router *gin.Engine) {
	
	// 发送验证码
	router.POST("/sendverifycode", service.SendVerifyCode)

	// 用户注册
	router.POST("/register", service.Register)

	// 用户登录
	router.POST("/login", service.Login)


	// 路由分组
	auth := router.Group("/u")
	auth.Use(middleware.AuthCheck())

	// 用户详情
	auth.GET("/user/detail", service.UserDetail)

	// 发送、接受消息
	auth.GET("/websocket/message", service.WebsocketMessage)


}
