package handlers

import (
	"Athena/internal/hub"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
)

var MainHub *hub.Hub

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print(err)
		return
	}

	// Creating a new client
	client := &hub.Client{
		ID:        uuid.New().String(),
		Conn:      conn,
		Send:      make(chan []byte, 256),
		SessionID: "default",
	}

	MainHub.Register <- client

	// Using readPump and writePump for modularity
	go readPump(client)
	go writePump(client)
}

func readPump(client *hub.Client) {
	// Autocloses after ReadPump is Done
	defer func() {
		MainHub.Unregister <- client
		client.Conn.Close()
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			break
		}
		MainHub.Broadcast <- message
	}
}

func writePump(client *hub.Client) {
	defer client.Conn.Close()
	for message := range client.Send {
		client.Conn.WriteMessage(websocket.TextMessage, message)
	}
}
