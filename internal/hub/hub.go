package hub

// Gives a Pointer to new Hub
func NewHub() *Hub {
	h := &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
	go h.run()
	return h
}

/*
Handles client registration, unregistration, and broadcasting.
Infinite for loop to handle multiple, simultaneous, and ongoing events reliably.

	Register - "Hey, new client wants in"
	Unregister - "Client's leaving, clean up"
	Broadcast - "Send this message to everyone"
*/
//The Instance built using NewHub() is given as a Receiver type to run()
func (h *Hub) run() {

	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
