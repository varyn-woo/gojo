package handlers

import (
	"gojo/gen"
	"log"
)

func RouteUserInput(input *gen.UserInputRequest) error {
	switch t := input.Request.(type) {
	case *gen.UserInputRequest_PlayerAddRequest:
		log.Printf("Received player add request: %s with display name: %s", t.PlayerAddRequest.PlayerId, t.PlayerAddRequest.DisplayName)
		return HandleAddPlayer(t.PlayerAddRequest)
	case *gen.UserInputRequest_StartGameRequest:
		log.Println("Received start game request")
		return HandleStartGame()
	default:
		return HandleUserInput(input)
	}
}
