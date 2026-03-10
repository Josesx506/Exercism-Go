package websockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Console struct {
	name string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: false,
}

func (c *Console) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "error accepting websocket connection", 400)
		return
	}
	// ... Use conn to send and receive messages.
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	if err := conn.WriteMessage(messageType, p); err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
}
