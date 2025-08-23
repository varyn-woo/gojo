package state

import (
	"errors"
	api_types "gojo/types"
	"sync"
)

type Game struct {
	Name    string
	Players []api_types.Player
	lock    sync.RWMutex
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

func (g *Game) GetGameObj() api_types.GameStateResponse {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return api_types.GameStateResponse{
		Name:    g.Name,
		Players: g.Players,
	}
}

func (g *Game) AddPlayer(player api_types.Player) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.Players = append(g.Players, player)
}

func (g *Game) RemovePlayer(player api_types.Player) {
	g.lock.Lock()
	defer g.lock.Unlock()
	for i, p := range g.Players {
		if p.Name == player.Name {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
		}
	}
}

func (g *Game) ChangeGame(name string) {
	g.Name = name
}

func (g *Game) GetPlayers() []api_types.Player {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.Players
}
