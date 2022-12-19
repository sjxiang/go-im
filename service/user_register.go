package service

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/go-im/helper"
	"github.com/sjxiang/go-im/model"
)

func Register(ctx *gin.Context) {
	verifyCode := ctx.PostForm("verify_code")
	email := ctx.PostForm("email")
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	if verifyCode == "" || email == "" || account == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "参数为空",
		})
		return
	}

	// 校验验证码
	r, err := model.RDB.Get(context.Background(), "Token_" + email).Result()
	if err != nil {
		log.Printf("[cache Error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "Redis 读取错误" + err.Error(),
		})

		return
	}

	if r != verifyCode {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "验证码不正确",
		})

		return
	}

	// 判断账号是否唯一
	cnt, err := model.GetUserBasicCountByAccount(account)
	if err != nil {
		log.Printf("[DB Error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "系统错误" + err.Error(),
		})

		return
	}

	if cnt > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "账号已被注册",
		})

		return
	}

	// 保存
	ub := &model.UserBasic{
		Identity: helper.GetUUID(),
		Account: account,
		Email: email,
		Password: helper.GetMd5(password),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	
	err = model.InsertOneUserBasic(ub)
	if err != nil {
		log.Printf("[DB Error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "系统错误" + err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200,
		"Msg": "注册成功",
	})
}