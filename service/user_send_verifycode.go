package service


import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/go-im/helper"
	"github.com/sjxiang/go-im/model"
)


// 发送验证码
func SendVerifyCode(ctx *gin.Context) {
	email := ctx.PostForm("email")

	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "邮箱不能为空",
		})

		return
	}

	
	cnt, err := model.GetUserBasicByEmail(email)
	if err != nil {
		log.Printf("[DB error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "数据库查询异常",
		})

		return
	}

	if cnt > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "当前邮箱已经被注册",
		})

		return
	}

	// 获取验证码
	verifyCode := helper.GetRandomNum()

	// TODO：redis 保留验证码
	

	// 发送验证码
	err = helper.NewEmail().Send(email, verifyCode)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "系统错误" + err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200,
		"Msg": "验证码发送成功",
	})
}