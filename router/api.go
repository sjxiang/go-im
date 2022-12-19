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
	v1 := router.Group("/v1")
	v1.Use(middleware.AuthCheck())

	// 用户详情
	v1.GET("/user/detail", service.UserDetail)

	// 创建聊天室
	v1.POST("/room/create", service.CreateRoom)

	// 发送、接受消息
	v1.GET("/websocket/message", service.WebsocketMessage)



	v2 := router.Group("/v2", middleware.AuthCheck())
	v2.GET("chat", service.WebsocketMessagePlus)
}
