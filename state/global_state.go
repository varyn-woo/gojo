package state

import "errors"

type Game struct {
	Name    string
	Players map[string]*Player
}

var game *Game

func NewGame() *Game {
	newGame := Game{
		Players: make(map[string]*Player),
	}
	game = &newGame
	return game
}

func GetGame() (*Game, error) {
	if game == nil {
		return nil, errors.New("game not initialized")
	}
	return game, nil
}

func (g *Game) AddPlayer(player *Player) error {
	if _, ok := g.Players[player.Name]; ok {
		return errors.New("player already exists")
	}
	g.Players[player.Name] = player
	return nil
}

func (g *Game) RemovePlayer(player *Player) {
	g.Players[player.Name] = nil
}

func (g *Game) ChangeGame(name string) {
	g.Name = name
}
