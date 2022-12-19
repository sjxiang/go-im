package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sjxiang/go-im/helper"
	"github.com/sjxiang/go-im/model"
)

type Message struct {
	Data string `json:"data"`
	RoomIdentity string `json:"room_identity"`
}

var upgrader = websocket.Upgrader{}  // http 替换成 websocket
var wc = make(map[string]*websocket.Conn)


func WebsocketMessage(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "系统错误" + err.Error(),
		})
		return
	}
	defer conn.Close()

	uc := ctx.MustGet("user_claims").(*helper.UserClaims)
	wc[uc.Identity] = conn

	for {
		ms := new(Message)
		
		err := conn.ReadJSON(ms)
		if err != nil {
			log.Printf("[Websocket Read Message Error]: %v\n", err)
			return
		}
		
		// TODO: 判断用户是否属于消息体的房间，校验
		_, err = model.GetUserRoomByUserIdentityRoomIdentity(uc.Identity, ms.RoomIdentity)
		if err != nil {
			log.Printf("[Websocket UserIdentity: %v RoomIdentity: %v Not Exists]\n", uc.Identity, ms.RoomIdentity)
			return
		}

		// TODO: 保存一份消息

		
		// TODO: 获取在特定房间的在线用户，消息推送

		for _, cc := range wc {
			err := cc.WriteMessage(websocket.TextMessage, []byte(ms.Data))
			if err != nil {
				log.Printf("[Websocket Read Message Error]: %v\n", err)
				return
			}
		}
	}
}