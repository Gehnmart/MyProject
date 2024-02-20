package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
	Content  string `json:"content"`
	RoomId   uint64 `json:"room_id"`
	Username string `json:"username"`
}

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	Id       uint64
	RoomId   uint64
	Username string
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		msg, ok := <-c.Message
		if !ok {
			return
		}

		err := c.Conn.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content:  string(m),
			RoomId:   c.RoomId,
			Username: c.Username,
		}

		hub.Broadcast <- msg
	}
}
