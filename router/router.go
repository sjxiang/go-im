package router

import (
	"github.com/gin-gonic/gin"

	"strings"
	"net/http"
)


func Setup() *gin.Engine {
	r := gin.Default()
	
	// 注册业务路由
	registerApiRoutes(r)

	// 注册全局中间件
	registerGlobalMiddleWare(r)

	// 配置 404 路由
	setupNoFoundHandler(r)


	return r
}


func registerGlobalMiddleWare(router *gin.Engine) {
	
}


func setupNoFoundHandler(router *gin.Engine) {

	// 处理错误路由，精确匹配
	router.NoRoute(func(ctx *gin.Context) {

		// 获取 header 里面的 'Accept' 信息
		acceptStr := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptStr, "text/html") {

			// 如果是 HTML
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {

			// 默认返回 JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确",
			})
		}
	})
}