package handlers

import (
	"gojo/gen"
	"gojo/state"
)

func HandleAddPlayer(req *gen.PlayerAddRequest) *gen.ServerResponse {
	game, err := state.GetGame()
	if err != nil {
		return MakeErrorResponse(err)
	}
	game.AddPlayer(&gen.Player{
		Id:             req.PlayerId,
		DisplayName:    req.DisplayName,
		IsActive:       true,
		IsPendingInput: false,
	})
	return MakeAcknowledgementResponse()
}
