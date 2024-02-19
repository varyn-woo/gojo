package state

import "errors"

type Game struct {
	Name    string
	Players []Player
}

var game *Game

func NewGame() *Game {
	newGame := Game{}
	game = &newGame
	return game
}

func GetGame() (*Game, error) {
	if game == nil {
		return nil, errors.New("game not initialized")
	}
	return game, nil
}

func (g *Game) AddPlayer(player Player) {
	g.Players = append(g.Players, player)
}

func (g *Game) RemovePlayer(player Player) {
	for i, p := range g.Players {
		if p.Name == player.Name {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
		}
	}
}

func (g *Game) ChangeGame(name string) {
	g.Name = name
}
