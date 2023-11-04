package api

import (
	"net/http"
	"operations/models"
	"strings"

	"github.com/gin-gonic/gin"
)

// getListaCantores responde com a lista de todos os álbuns em JSON.
func getListaCantores(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, GetCantores())
}

// postAlbums adiciona um álbum a partir do JSON recebido no corpo da requisição.
func postCantores(c *gin.Context) {
	var newCantor models.Cantor

	// Chama BindJSON para vincular o JSON recebido a newCantor.
	if err := c.BindJSON(&newCantor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifique se os campos necessários não são vazios
	if newCantor.Nome == "" || newCantor.Mesa == "" || newCantor.CodigoMusica == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos os campos são obrigatórios"})
		return
	}

	AddCantor(newCantor)
	c.IndentedJSON(http.StatusCreated, newCantor)
}

// Inicialização da lista de clientes
var cantores = []models.Cantor{}

// GetCantores retorna a lista de álbuns.
func GetCantores() []models.Cantor {
	return cantores
}

// AddCantor adiciona um álbum à lista.
func AddCantor(newCantor models.Cantor) {
	cantores = append(cantores, newCantor)
}

func removeItem(c *gin.Context) {
	// Obter o ID do item a ser removido a partir dos parâmetros da URL
	id := strings.TrimSpace(c.Param("nome"))

	// Percorre o slice de itens para encontrar e remover o item com o ID correspondente
	indexToRemove := -1
	for i, item := range cantores {
		if strings.TrimSpace(item.Nome) == id {
			indexToRemove = i
			break
		}
	}

	// Se o item com o ID correspondente foi encontrado, remova-o
	if indexToRemove != -1 {
		cantores = append(cantores[:indexToRemove], cantores[indexToRemove+1:]...)
		c.JSON(http.StatusOK, gin.H{"message": "Item removido com sucesso"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item não encontrado"})
	}
}
