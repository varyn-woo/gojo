package handlers

import (
	"errors"
	"gojo/gen"
	"log"
)

func RouteUserInput(input *gen.UserInputRequest) *gen.ServerResponse {
	switch t := input.Request.(type) {
	case *gen.UserInputRequest_PlayerAddRequest:
		log.Printf("Received player add request: %s with display name: %s", t.PlayerAddRequest.PlayerId, t.PlayerAddRequest.DisplayName)
		return HandleAddPlayer(t.PlayerAddRequest)
	case *gen.UserInputRequest_StartGameRequest:
		log.Println("Received start game request")
		return HandleStartGame()
	case *gen.UserInputRequest_NewGameRequest:
		log.Println("Received new game request")
		return HandleNewGame()
	default:
		log.Printf("Received unknown input type: %T", t)
		return MakeErrorResponse(errors.New("not implemented"))
	}
}
