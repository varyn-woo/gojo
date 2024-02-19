package handlers

import (
	"gojo/state"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleSetGame(c *gin.Context) {
	game, err := state.GetGame()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var gameConfig state.GameConfig
	c.BindJSON(&gameConfig)
	log.Printf("changing game to %s", gameConfig.Game)
	game.ChangeGame(gameConfig.Game)
	c.JSON(200, game)
}

func HandleGetGame(c *gin.Context) {
	game, err := state.GetGame()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, game)
}

func HandleAddPlayer(c *gin.Context) {
	game, err := state.GetGame()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var player state.PlayerConfig
	c.BindJSON(&player)
	err = game.AddPlayer(&state.Player{Name: player.Name})
	if err != nil {
		log.Printf("adding player: %s failed: %v", player.Name, err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.Printf("added player: %s", player.Name)
	c.JSON(200, game)
}
