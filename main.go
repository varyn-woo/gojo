package main

import (
	"fmt"
	"gojo/gserver"
	"gojo/state"
)

const (
	serverIP   = "localhost"
	serverPort = "8080"
)

func main() {
	// initialize the game
	state.NewGame()
	r := gserver.GetGameServer()
	// listen and serve on port :8080
	r.Run(fmt.Sprintf("%s:%s", serverIP, serverPort))
}
