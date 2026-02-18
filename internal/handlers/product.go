package handlers

import (
	"net/http"
	"portal/internal/models"
	"portal/internal/services"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Produtos Registrados",
		"data":    services.ListProducts(),
	})
}

func GetProductById(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Produto encotrado pelo ID: " + id,
		"data":    services.GetProductById(id),
	})
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	p, err := services.CreateProduct(product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Produdo criado com sucesso",
		"data":    p,
	})
}

func UpdateProduct(c *gin.Context) {

	var p models.Product
	id := c.Query("id")

	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	p, err := services.UpdateProduct(id, p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Produdo atualizado com sucesso",
		"data":    p,
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Query("id")

	err := services.DeleteProduct(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Produdo deletado com sucesso",
	})
}
