package state

import (
	"errors"
	"gojo/gen"
	"sync"
)

type Game struct {
	started bool
	state   *gen.GameState
	lock    sync.RWMutex
}

var game *Game

func NewGame() *Game {
	newGame := Game{
		state: &gen.GameState{
			CurrentStage: gen.GameStage_AWAITING_PLAYERS,
		},
	}
	game = &newGame
	return game
}

func GetGame() (*Game, error) {
	if game == nil {
		return NewGame(), nil
	}
	return game, nil
}

func (g *Game) StartGame() {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.started {
		return // game already started
	}
	g.started = true
	g.state.CurrentStage = gen.GameStage_COLLECTING_RESPONSES
}

func (g *Game) GetGameState() *gen.GameState {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.state
}

func (g *Game) AddPlayer(player *gen.Player) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if len(g.state.Players) == 0 {
		g.state.HostPlayerId = player.Id // set the first player as the host
	}
	g.state.Players = append(g.state.Players, player)
}

func (g *Game) RemovePlayer(player *gen.Player) {
	g.lock.Lock()
	defer g.lock.Unlock()
	i := g.getPlayerIndexLocked(player.Id)
	if i >= 0 {
		g.state.Players = append(g.state.Players[:i], g.state.Players[i+1:]...)
		return
	}
}

func (g *Game) UpdatePlayer(player *gen.Player) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	i := g.getPlayerIndexLocked(player.Id)
	if i < 0 {
		return errors.New("player not found")
	}
	g.state.Players[i] = player
	return nil
}

// getPlayerIndexLocked gets the index of a player base on their ID.
// It assumes an appropriate lock has already been taken.
func (g *Game) getPlayerIndexLocked(playerId string) int {
	for i, player := range g.state.Players {
		if player.Id == playerId {
			return i
		}
	}
	return -1
}

func (g *Game) GetPlayers() []*gen.Player {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.state.Players
}
