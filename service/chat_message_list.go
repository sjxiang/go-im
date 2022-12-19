package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/go-im/helper"
	"github.com/sjxiang/go-im/model"
)

func ChatMessageList(ctx *gin.Context) {
	
	pageIndex := ctx.Query("page_index")  // 第 x 页
	pageSize := ctx.Query("page_size")    // 每页 y 条记录
	roomIdentity := ctx.Query("room_identity")

	if roomIdentity == "" || pageIndex == "" || pageSize == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "参数不能为空",
		})

		return
	}

	// 判断用户是否属于该房间
	uc := ctx.MustGet("user_claims").(*helper.UserClaims)
	
	_, err := model.GetUserRoomByUserIdentityRoomIdentity(uc.Identity, roomIdentity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "非法访问",
		})
		return
	}

	// 聊天记录查询
	pageIndexNum, err := strconv.ParseInt(pageIndex, 10, 32)
	if err != nil {
		log.Printf("[请求参数错误]:%v\n", err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "请求参数错误" + err.Error(),
		})
		return
	}

	pageSizeNum, err := strconv.ParseInt(pageSize, 10, 32)
	if err != nil {
		log.Printf("[请求参数错误]:%v\n", err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "请求参数错误" + err.Error(),
		})
		return
	}

	skip := (pageIndexNum - 1) * pageSizeNum
	data, err := model.GetMessageListByRoomIdentity(roomIdentity, pageSizeNum, skip)
	if err != nil {
		log.Printf("[DB Error]: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "系统异常" + err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200,
		"Msg": "数据加载成功",
		"Data": gin.H{
			"message_list": data,
		},
	})
}