package handlers

import (
	"gojo/gen"
	"gojo/state"
)

func HandleAddPlayer(req *gen.PlayerAddRequest) error {
	game, err := state.GetGame()
	if err != nil {
		return err
	}
	game.AddPlayer(&gen.Player{
		Id:          req.PlayerId,
		DisplayName: req.DisplayName,
		IsActive:    true,
		IsWaiting:   false,
	})
	return nil
}
