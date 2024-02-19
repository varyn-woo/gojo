package state

type Player struct {
	Name   string
	Points int
}

type GameConfig struct {
	Game string `json:"game"`
}
