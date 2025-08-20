package gserver

import (
	"gojo/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetGameServer() *gin.Engine {
	r := gin.Default()

	// handle CORS (pre-flight HTTP request authorization)
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// set handlers
	r.GET("/game_get", handlers.HandleGetGame)
	r.POST("/game_set", handlers.HandleSetGame)
	r.POST("/psych_input", handlers.HandleAddPsychTextInput)
	r.POST("/player_add", handlers.HandleAddPlayer)
	r.POST("/player_remove", handlers.HandleRemovePlayer)
	r.GET("/players_get", handlers.HandleGetPlayers)

	return r
}
