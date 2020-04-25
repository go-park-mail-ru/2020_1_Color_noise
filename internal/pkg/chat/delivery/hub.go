package delivery


import "2020_1_Color_noise/internal/models"

type Hub struct {
	// Registered clients.
	clients map[uint]*Client

	// Inbound messages from the clients.
	broadcast chan *models.Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *models.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[uint]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.userId] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.userId]; ok {
				delete(h.clients, client.userId)
				close(client.send)
			}
		case message := <-h.broadcast:
			client := h.clients[message.RecUser.Id]
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client.userId)
			}
		}
	}
}
