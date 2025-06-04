package routes

import (
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
