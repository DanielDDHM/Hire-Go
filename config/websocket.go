package server

import (
	"log"
	"net/http"

	"github.com/DanielDDHM/Hire-Go/internal/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

var broadcast = make(chan models.Message)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erro ao abrir conexão WebSocket:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Erro ao ler mensagem:", err)
			delete(clients, conn)
			break
		}

		broadcast <- msg
	}
}

func StartBroadcaster() {
	for {
		msg := <-broadcast
		for conn := range clients {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println("Erro ao enviar mensagem:", err)
				conn.Close()
				delete(clients, conn)
			}
		}
	}
}
