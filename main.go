package main

import (
	"operations/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configuração das rotas da API
	api.SetupRoutes(router)

	router.Run("localhost:8080")
}
