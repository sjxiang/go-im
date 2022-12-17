package service

import (
	"log"
	"net/http"

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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "用户名或密码错误"+err.Error(),
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


// 用户详情
func UserDetail(ctx *gin.Context) {
	u, _ := ctx.Get("user_claims")
	uc := u.(*helper.UserClaims)

	userBasic, err := model.GetUserBasicByIdentity(uc.Identity)
	if err != nil {
		log.Printf("[DB error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "数据库查询异常",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200,
		"Msg": "数据加载成功",
		"Data": userBasic,
	})
}


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
	verifyCode := helper.GenerateRandomNum()

	// TODO：redis 保留一份验证码
	

	// 邮箱发送验证码
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