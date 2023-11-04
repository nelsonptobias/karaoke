package main

import (
	"operations/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configuração das rotas da API
	api.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Use a porta padrão se a variável de ambiente não estiver definida
	}

	router.Run(":" + port)
}
