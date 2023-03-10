package service



import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/go-im/helper"
	"github.com/sjxiang/go-im/model"
)


// 用户登录
func Login(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	

	if account == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "用户名或密码不能为空",
		})
		return
	}

	ub, err := model.GetUserBasicByUsernamePassword(account, helper.GetMd5(password))
	if err != nil {
		log.Printf("[DB Error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "数据查询异常" + err.Error(),  // 也有可能账号、密码输入错误
		})
		return
	}

	token, err := helper.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "系统错误" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200,
		"Msg": "登录成功",
		"Data": gin.H{
			"token": token,
		},
	})
}
