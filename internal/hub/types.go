package hub

import (
	"github.com/gorilla/websocket"
)

// A single WebSocket connection.
type Client struct {
	ID        string
	Conn      *websocket.Conn
	Send      chan []byte
	SessionID string
}

// Hub manages all active clients and message routing.
type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}
