package handlers

import (
	"gojo/state"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleAddPlayer(c *gin.Context) {
	game, err := state.GetGame()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var player state.Player
	c.BindJSON(&player)
	log.Printf("adding player: %s", player.Name)
	game.AddPlayer(player)
	c.JSON(200, game)
}

func HandleRemovePlayer(c *gin.Context) {
	game, err := state.GetGame()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var player state.Player
	c.BindJSON(&player)
	log.Printf("removing player: %s", player.Name)
	game.RemovePlayer(player)
	c.JSON(200, game)
}

func HandleGetPlayers(c *gin.Context) {
	game, err := state.GetGame()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	players := game.GetPlayers()
	c.JSON(200, players)
}
