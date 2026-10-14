package websocket

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"realtime-server/internal/chat"
	"realtime-server/internal/common/db"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room")
	userID := r.URL.Query().Get("user")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	client := &Client{
		Conn:   conn,
		UserID: userID,
		RoomID: roomID,
		Send:   make(chan []byte),
	}

	WS.Register <- client

	go client.readPump()
	go client.writePump()
}

func (c *Client) readPump() {
	chatRepo := chat.NewRepository(db.InitMongo())
	defer func() {
		WS.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		var m chat.Message
		json.Unmarshal(msg, &m)
		m.SenderID, _ = bson.ObjectIDFromHex(c.UserID)
		m.RoomID = c.RoomID

		chatRepo.SaveMessage(context.TODO(), &m)

		WS.Broadcast <- MessagePayload{
			RoomID:  c.RoomID,
			Message: msg,
		}
	}
}

func (c *Client) writePump() {
	for msg := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
