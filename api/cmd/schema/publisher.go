package schema

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/graphql-go/graphql"
	"github.com/jjg-akers/docker-sql-graphql/cmd/subscriptions"
)

func PublicationPublisher() {
	subscriptions.Subscribers.Range(func(key, value interface{}) bool {
		subscriber, ok := value.(*subscriptions.Subscriber)
		if !ok {
			return true
		}

		payload := graphql.Do(graphql.Params{
			Schema:        Schema,
			RequestString: subscriber.RequestStr,
		})

		message, err := json.Marshal(map[string]interface{}{
			"type":    "data",
			"id":      subscriber.OperationID,
			"payload": payload,
		})
		if err != nil {
			log.Println("failed to marshal message: ", err)
		}
		if err := subscriber.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			if err == websocket.ErrCloseSent {
				subscriptions.Subscribers.Delete(key)
				return true
			}
			log.Println("failed to write to we connection: ", err)
			return true
		}
		return true
	})
}
