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
	log.Printf("Changing game to %s", gameConfig.Game)
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
