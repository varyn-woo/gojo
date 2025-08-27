package state

import (
	"errors"
	"gojo/gen"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type GameHandler interface {
	Sync() error
	HandleUserInput(*gen.UserInputRequest) error
}

type Game struct {
	players     map[string]*gen.Player
	gameHandler GameHandler
	state       *gen.GameState
	lock        sync.RWMutex
}

type StateUpdater func(gs *gen.GameState)

var game *Game

var (
	ErrOutOfSync    = errors.New("input out of sync")
	ErrInvalidInput = errors.New("invalid input type for this game")
)

func NewGame() *Game {
	newGame := Game{
		gameHandler: &PsychHandler{timerDuration: DEFAULT_TIMER},
		state:       &gen.GameState{},
		players:     map[string]*gen.Player{},
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
	if g.state.Started {
		g.lock.Unlock()
		return // game already started
	}
	g.state.Started = true
	g.state.TimerStart = timestamppb.Now()
	g.lock.Unlock()
	g.gameHandler.Sync()
}

func (g *Game) ResetTimer() {
	g.state.TimerStart = timestamppb.Now()
}

func (g *Game) GetElapsedTime() time.Duration {
	return time.Since(g.state.TimerStart.AsTime())
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

	g.players[player.Id] = player
	g.syncPlayerListLocked()
}

func (g *Game) RemovePlayer(player *gen.Player) {
	g.lock.Lock()
	defer g.lock.Unlock()
	delete(g.players, player.Id)
	g.syncPlayerListLocked()
}

func (g *Game) IncrPlayerScore(id string, score int) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	player, exists := g.players[id]
	if !exists {
		return errors.New("player not found")
	}
	player.Score += int32(score)
	g.syncPlayerListLocked()
	return nil
}

func (g *Game) SetPlayerWaiting(id string, isWaiting bool) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	player, exists := g.players[id]
	if !exists {
		return errors.New("player not found")
	}
	player.IsWaiting = isWaiting
	g.syncPlayerListLocked()
	return nil
}

func (g *Game) EndWait() {
	g.lock.Lock()
	defer g.lock.Unlock()
	for _, p := range g.players {
		p.IsWaiting = false
	}
	g.syncPlayerListLocked()
}

// syncPlayerListLocked updates the players array in the game state
// to reflect the actual players. Call it after any player update
func (g *Game) syncPlayerListLocked() {
	playerList := []*gen.Player{}
	for _, player := range g.players {
		playerList = append(playerList, player)
	}
	game.state.Players = playerList
}

func (g *Game) GetPlayers() []*gen.Player {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.state.Players
}

// UpdateState uses a mutator function to transform the state.
// This allows for generic state updates that don't have race conditions.
func (g *Game) HandleInput(userInput *gen.UserInputRequest) error {
	err := g.gameHandler.HandleUserInput(userInput)
	if err != nil {
		return err
	}
	err = g.gameHandler.Sync()
	if err != nil {
		return err
	}
	return nil
}
