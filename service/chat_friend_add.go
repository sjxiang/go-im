package service

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/go-im/helper"
	"github.com/sjxiang/go-im/model"
)


func FriendAdd(ctx *gin.Context) {
	account := ctx.PostForm("account")
	
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

	
	if ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "互为好友，不可重复添加",
		})

		return
	}

	// 聊天室
	rb := &model.RoomBasic{
		Identity: helper.GetUUID(),
		UserIdentity: uc.Identity,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	err = model.InsertOneRoomBasic(rb)
	if err != nil {
		log.Printf("[DB Error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "数据插入异常" + err.Error(),
		})

		return
	}


	// 保存用户-聊天室 关联记录
	ur := &model.UserRoom{
		UserIdentity: uc.Identity,
		RoomIdentity: rb.Identity,
		RoomType: 1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	err = model.InsertOneUserRoom(ur)
	if err != nil {
		log.Printf("[DB Error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "数据插入异常" + err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200,
		"Msg": "添加好友成功",
	})
	
}