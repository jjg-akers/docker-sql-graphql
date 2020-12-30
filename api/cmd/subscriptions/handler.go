package subscriptions

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type ConnACKMessage struct {
	OperationID string `json:"id,omitempty"`
	Type        string `json:"type"`
	Payload     struct {
		Query string `json:"query"`
	} `json:"payload,omitempty"`
}

type Subscriber struct {
	ID          int
	Conn        *websocket.Conn
	RequestStr  string
	OperationID string
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Subprotocols: []string{"graphql-ws"},
}

var Subscribers sync.Map

//"Unexpected token B in JSON at position 0"

//Handler accepts the incoming websocket request
func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to upgrade websocket: %v\n", err)
		return
	}

	connACK, err := json.Marshal(map[string]string{
		"type": "connection_ack",
	})
	if err != nil {
		log.Println("Failed to marshal ws conn ack: ", err)
	}

	if err := conn.WriteMessage(websocket.TextMessage, connACK); err != nil {
		log.Println("failed to write to ws connection: ", err)
		return
	}

	//  listen for messages from the socket
	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				return
			}
			if err != nil {
				log.Println("failed to read websocket message: ", err)
			}

			var msg ConnACKMessage
			if err := json.Unmarshal(p, &msg); err != nil {
				log.Println("failed to unmarshal: ", err)
				return
			}

			if msg.Type == "start" {
				length := 0
				Subscribers.Range(func(key, value interface{}) bool {
					length++
					return true
				})

				subscriber := Subscriber{
					ID:          length + 1,
					Conn:        conn,
					RequestStr:  msg.Payload.Query,
					OperationID: msg.OperationID,
				}

				// list of active subscribers that will be notified depending on the graphql subscirption Query they pass
				Subscribers.Store(subscriber.ID, &subscriber)
			}
		}
	}()
}
