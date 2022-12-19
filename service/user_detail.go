package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/go-im/helper"
	"github.com/sjxiang/go-im/model"
)

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
		"Data": userBasic,  // TODO 序列化器，过滤敏感信息
	})
}

