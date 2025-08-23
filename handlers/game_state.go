package handlers

import (
	"gojo/gen"
	"gojo/state"
)

func HandleNewGame() *gen.ServerResponse {
	game := state.NewGame()
	return &gen.ServerResponse{
		Response: &gen.ServerResponse_GameState{
			GameState: game.GetGameState(),
		},
	}
}

func HandleStartGame() *gen.ServerResponse {
	game, err := state.GetGame()
	if err != nil {
		return MakeErrorResponse(err)
	}
	game.StartGame()
	return MakeGameStateResponse()
}
