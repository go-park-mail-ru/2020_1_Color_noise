package delivery

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/chat"
	e "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/response"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	userId uint

	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan *models.Message

	usecase chat.IUsecase
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		input := &models.InputMessage{}

		err := c.conn.ReadJSON(input)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message, err := c.usecase.AddMessage(c.userId, input)
		if err != nil {
			log.Println(err)
			break
		}

		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			resp := &models.ResponseMessage{
				SendUser: &models.ResponseUser{
					Id:            message.SendUser.Id,
					Login:         message.SendUser.Login,
					About:         message.SendUser.About,
					Avatar:        message.SendUser.Avatar,
					Subscribers:   message.SendUser.Subscribers,
					Subscriptions: message.SendUser.Subscriptions,
				},

				RecUser: &models.ResponseUser{
					Id:            message.SendUser.Id,
					Login:         message.SendUser.Login,
					About:         message.SendUser.About,
					Avatar:        message.SendUser.Avatar,
					Subscribers:   message.SendUser.Subscribers,
					Subscriptions: message.SendUser.Subscriptions,
				},

				Message: message.Message,

				Stickers: message.Stickers,

				CreatedAt: message.CreatedAt,
			}

			result := response.Response{
				Status: 200,
				Body:   resp,
			}

			err := c.conn.WriteJSON(result)
			if err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, logger *zap.SugaredLogger, usecase chat.IUsecase, w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		//log.Println("User is unauthorized, reqId: ", reqId)
		err := e.Unauthorized.New("Chatting: user is unauthorized")
		e.ErrorHandler(w, r, logger, reqId, e.Wrapf(err, "request id: %s", reqId))
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := e.NoType.New("Received bad id from context")
		e.ErrorHandler(w, r, logger, reqId, e.Wrapf(err, "request id: %s", reqId))
		return
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		err := e.NoType.New("Check origin error, chat")
		e.ErrorHandler(w, r, logger, reqId, e.Wrapf(err, "request id: %s", reqId))
		return
	}

	client := &Client{userId: userId, hub: hub, conn: conn, send: make(chan *models.Message), usecase: usecase}
	client.hub.register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
