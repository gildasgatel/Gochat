package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
)

type Handler interface {
	Home(c *gin.Context)
	WS(c *gin.Context)
}

type Handle struct {
}

func New() Handler {
	return &Handle{}
}

func (h *Handle) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
func (h *Handle) WS(c *gin.Context) {
	handleWebSocket(c.Writer, c.Request)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		broadcastMessage(msg)
	}
}

func broadcastMessage(msg []byte) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println(err)
			client.Close()
			delete(clients, client)
		}
	}
}
