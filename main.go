package main

import (
	"fmt"
	"log"
	"net/http"

	"gojo/gen"
	"gojo/handlers"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
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
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Error %s when reading message from client", err)
			return
		}
		var input gen.UserInputRequest
		err = proto.Unmarshal(message, &input)
		if err != nil {
			log.Printf("Error %s when unmarshalling message from client", err)
			err = c.WriteMessage(websocket.TextMessage, []byte("Invalid input format"))
			handleWriteErr(err)
			continue
		}

		// send user input to input handlers to get a response obj
		resp := handlers.RouteUserInput(&input)
		respBytes, err := proto.Marshal(resp)
		if err != nil {
			log.Printf("Error %s when marshalling response to client", err)
			err = c.WriteMessage(websocket.TextMessage, []byte("Error processing request"))
			handleWriteErr(err)
			continue
		}
		// send successful response obj
		err = c.WriteMessage(websocket.BinaryMessage, respBytes)
		handleWriteErr(err)
	}

}

func main() {
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(_ *http.Request) bool { return true },
		},
	}

	http.Handle("/", webSocketHandler)
	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", serverIP, serverPort), nil))
}
