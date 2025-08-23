package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gojo/gen"
	"gojo/handlers"

	"github.com/gorilla/websocket"
)

const (
	serverIP   = "localhost"
	serverPort = "8080"
)

// func main() {
// 	// initialize the game
// 	state.NewGame()
// 	r := gserver.GetGameServer()
// 	// listen and serve on port :8080
// 	r.Run(fmt.Sprintf("%s:%s", serverIP, serverPort))
// }

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}

	handleWriteErr := func(err error) {
		if err != nil {
			log.Printf("error %s when writing to websocket client", err)
			return
		}
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Error %s when reading message from client", err)
			return
		}
		if mt == websocket.BinaryMessage {
			err = c.WriteMessage(websocket.TextMessage, []byte("server doesn't support binary messages"))
			handleWriteErr(err)
			return
		}
		log.Printf("Receive message %s", string(message))
		var input gen.UserInput
		err = json.Unmarshal(message, &input)
		if err != nil {
			log.Printf("Error %s when unmarshalling message from client", err)
			err = c.WriteMessage(websocket.TextMessage, []byte("Invalid input format"))
			handleWriteErr(err)
			continue
		}

		// send user input to input handlers to get a response obj
		resp := handlers.RouteUserInput(&input)
		respJson, err := json.Marshal(resp)
		if err != nil {
			log.Printf("Error %s when marshalling response to client", err)
			err = c.WriteMessage(websocket.TextMessage, []byte("Error processing request"))
			handleWriteErr(err)
		}
		// send successful response obj
		err = c.WriteMessage(websocket.TextMessage, respJson)
		handleWriteErr(err)
	}

}

func main() {
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}

	http.Handle("/", webSocketHandler)
	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", serverIP, serverPort), nil))

}
