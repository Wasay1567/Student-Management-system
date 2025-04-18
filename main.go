package main

import (
	"os"

	"github.com/Wasay1567/edutrack/config"
	"github.com/Wasay1567/edutrack/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()

	config.ConnectDB()
	routes.SetupRoutes(r)

	r.Run(os.Getenv("PORT"))

}
