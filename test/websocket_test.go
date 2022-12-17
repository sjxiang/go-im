package test

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)



var upgrader = websocket.Upgrader{}
var ws = make(map[*websocket.Conn]struct{})

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("read:" + err.Error())
		return
	}
	defer conn.Close()

	ws[conn] = struct{}{}

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:" + err.Error())
			break
		}

		log.Println("receive:", message)
		
		for conn := range ws {
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:" + err.Error())
				break
			}
		}
		
	}

}


func TestWebsocketServer(t *testing.T) {
	r := gin.Default()

	r.GET("/echo", func(ctx *gin.Context) {
		echo(ctx.Writer, ctx.Request)
	})

	r.Run(":8080")
}