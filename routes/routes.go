package routes

import (
	"effy/gravatar-profile-card/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func Routes(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://yourfrontenddomain.com"}, 
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},  
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, 
		AllowCredentials: true,  
	}))

	server.POST("/gravatardata", handlers.Gravatar)
}
