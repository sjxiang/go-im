package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/go-im/helper"
	"github.com/sjxiang/go-im/model"
)

type UserQueryResult struct {
	NickName string `json:"nickname"`
	Sex      int `json:"sex"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"is_friend"`
}


func UserQuery(ctx *gin.Context) {
	account := ctx.Query("account")
	
	if account == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "参数不能为空",
		})

		return
	}
	
	ub, err := model.GetUserBasicByAccount(account)
	if err != nil {
		log.Printf("[DB Error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "数据查询异常" + err.Error(),
		})

		return
	}

	uc := ctx.MustGet("user_claims").(*helper.UserClaims)

	ok, err := model.JudgeUserIsFriend(uc.Identity, ub.Identity); 
	if err != nil {
		log.Printf("[DB Error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "数据查询异常" + err.Error(),
		})

		return
	}

	
	data := UserQueryResult {
		NickName: ub.Nickname,
		Sex: ub.Sex,
		Email: ub.Email,
		Avatar: ub.Avatar,
		IsFriend: true,
	}

	if !ok {
		data.IsFriend = false
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200,
		"Msg": "登录成功",
		"Data": gin.H{
			"token": data,
		},
	})
	
}