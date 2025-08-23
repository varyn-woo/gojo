package handlers

import (
	"gojo/state"
	"io"

	"github.com/gin-gonic/gin"
)

func HandleSetGame(c *gin.Context) {
	game, err := state.GetGame()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	gameName, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid game name"})
		return
	}
	game.ChangeGame(string(gameName))
	c.JSON(200, game.GetGameObj())
}

func HandleGetGame(c *gin.Context) {
	game, err := state.GetGame()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, game.GetGameObj())
}
