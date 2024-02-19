package main

import (
	"gojo/handlers"
	"gojo/state"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize the game
	state.NewGame()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/game_get", handlers.HandleGetGame)
	r.POST("/game_set", handlers.HandleSetGame)
	// listen and serve on 0.0.0.0:8080
	// on windows "localhost:8080"
	// can be overriden with the PORT env var
	r.Run("localhost:8080")
}
