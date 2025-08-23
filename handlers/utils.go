package handlers

import (
	"gojo/gen"
	"gojo/state"
)

func MakeErrorResponse(err error) *gen.ServerResponse {
	if err == nil {
		return nil
	}
	return &gen.ServerResponse{
		Response: &gen.ServerResponse_ErrorMessage{
			ErrorMessage: err.Error(),
		},
	}
}

func MakeAcknowledgementResponse() *gen.ServerResponse {
	return &gen.ServerResponse{
		Response: &gen.ServerResponse_Acknowledgement{
			Acknowledgement: true,
		},
	}
}

func MakeGameStateResponse() *gen.ServerResponse {
	game, err := state.GetGame()
	if err != nil {
		return MakeErrorResponse(err)
	}
	return &gen.ServerResponse{
		Response: &gen.ServerResponse_GameState{
			GameState: game.GetGameState(),
		},
	}
}
