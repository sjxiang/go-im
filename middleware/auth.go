package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/go-im/helper"
)


// token 放在 Header 的 Authorization 中，例如 "bearer xxx.xxx.xxx"

func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg": "请求 header 中 auth 为空",
			})
			ctx.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusUnauthorized,  gin.H{
				"code": -1,
				"msg": "请求 header 中，auth 格式有误",
			})
			ctx.Abort()
			return
		}

		// parts[1] 为 token
		userClaims, err := helper.AnalyzeToken(parts[1]) 
		if err != nil {
			ctx.Abort()  // 提前中止请求处理
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Code": -1,
				"Msg": "用户认证不通过",
			})
			
			return
		}

		// 检查是否过期
		if float64(time.Now().Unix()) > float64(userClaims.ExpiresAt) {
			ctx.Abort()  // 提前中止请求处理
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Code": -1,
				"Msg": "token 过期",
			})
			
			return
		}



		ctx.Set("user_claims", userClaims)
		ctx.Next()
	}
}	