package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sjxiang/go-im/helper"
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

		for _, cc := range wc {
			err := cc.WriteMessage(websocket.TextMessage, []byte(ms.Data))
			if err != nil {
				log.Printf("[Websocket Read Message Error]: %v\n", err)
				return
			}
		}
	}
}