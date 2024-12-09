package routes

import (
	"effy/gravatar-profile-card/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func Routes(server *gin.Engine) {

	authHandler := handlers.NewAuthHandler()
	chandler := handlers.NewChandler()
	qrcode := handlers.NewQrCodehandler()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3000"}, 
		AllowMethods:     []string{"GET", "POST","PUT","PATCH","DELETE","OPTIONS"},  
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, 
		AllowCredentials: true,  
	}))

	server.POST("/gravatardata", handlers.Gravatar)
	server.POST("/upload",chandler.UploadAndGeneratePublicURL)
	server.POST("/qrcode",qrcode.GenerateQR)

	server.POST("/signup",authHandler.Signup)
	server.POST("/signin",authHandler.Signin)
}
