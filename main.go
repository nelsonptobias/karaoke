package main

import (
	"log"
	"operations/api"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configure o middleware CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "DELETE"}
	router.Use(cors.New(config))

	router.Use(gin.Logger())
	// Configuração das rotas da API
	api.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
		port = "8080" // Use a porta padrão se a variável de ambiente não estiver definida
	}

	router.Run(":" + port)
}
