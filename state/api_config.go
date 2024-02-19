package state

type Player struct {
	Name   string
	Points int
}

// GameConfig is a struct that all game configuration requests will be bound to
type GameConfig struct {
	Game string `json:"game"`
}

// PlayerConfig is a struct that all player configuration requests will be bound to
type PlayerConfig struct {
	Name string `json:"name"`
}
