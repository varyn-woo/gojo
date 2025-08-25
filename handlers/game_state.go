package handlers

import (
	"gojo/gen"
	"gojo/state"
)

func handleGameStateChange(updateFunc func(*state.Game)) error {
	game, err := state.GetGame()
	if err != nil {
		return err
	}
	updateFunc(game)
	return nil
}

func HandleNewGame() error {
	_ = state.NewGame()
	return nil
}

func HandleStartGame() error {
	return handleGameStateChange(func(g *state.Game) { g.StartGame() })
}

func HandleUserInput(input *gen.UserInputRequest) error {
	return handleGameStateChange(func(g *state.Game) {
		g.HandleInput(input)
	})
}
