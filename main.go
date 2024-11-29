package main

import (
	"effy/gravatar-profile-card/db"
	"effy/gravatar-profile-card/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitDB()
	if err != nil {
		panic("database error" + err.Error())
	}
	server := gin.Default()
	routes.Routes(server)
	server.Run("0.0.0.0:8080")

}
