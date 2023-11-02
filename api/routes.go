package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura as rotas da API.
func SetupRoutes(router *gin.Engine) {
	router.GET("/lista", getListaCantores)
	router.POST("/lista", postCantores)
	router.DELETE("/lista/:nome", removeItem)
}
