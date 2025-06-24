package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

const pongwait = 60 * time.Second

func mainPageHandler(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"Name": "Gin Framework"})
}

func gameHandler(ctx *gin.Context) {
	w, r := ctx.Writer, ctx.Request
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Smt wrong"})
		return
	}
	defer c.Close()
	var board GameBoard
	board.initiate()
	data, _ := json.Marshal(&Message{
		Type: "update_board",
		Data: map[string]any{"board": board},
	})
	c.SetWriteDeadline(time.Now().Add(3 * time.Second))
	err  = c.WriteMessage(1, data)
	if err != nil {
		log.Printf("Write msg error: %s", err)
		return
	}

	for {
		c.SetReadDeadline(time.Now().Add(pongwait))
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Read msg error: %s", err)
			saveGameResult(&board, ctx)
			break
		}
		log.Printf("received: %s", message)
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			res, _ := json.Marshal("Wrong message type")
			c.SetWriteDeadline(time.Now().Add(3 * time.Second))
			err = c.WriteMessage(mt, res)
			if err != nil {
				log.Printf("Write msg error: %s", err)
				break
			}
		}
		res := handleMessage(&msg, &board)
		if res.Type == "end_game"{
			saveGameResult(&board, ctx)
			break
		}
		ok := sendMessage(res, mt, c)
		if !ok {
			break
		}

	}
}

func sendMessage(msg *Message, mt int, c *websocket.Conn) bool {
	data, _ := json.Marshal(msg)
	c.SetWriteDeadline(time.Now().Add(3 * time.Second))
	err := c.WriteMessage(mt, data)
	if err != nil {
		log.Printf("Write msg error: %s", err)
		return false
	}
	return true
}