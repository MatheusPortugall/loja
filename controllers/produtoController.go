package controllers

import (
	"loja/database"
	"loja/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListaTodosProdutos(c *gin.Context) {
	var produtos []models.Produto
	database.DB.Find(&produtos)
	c.JSON(http.StatusOK, produtos)
}

func ObtemUmProduto(c *gin.Context) {
	id := c.Params.ByName("id")
	var produto models.Produto
	database.DB.First(&produto, id)
	c.JSON(http.StatusOK, produto)
}

func CriaNovoProduto(c *gin.Context) {
	var novoProduto models.Produto
	if err := c.ShouldBindJSON(&novoProduto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&novoProduto)
	c.JSON(http.StatusOK, novoProduto)
}

func AtualizarProduto(c *gin.Context) {
	id := c.Params.ByName("id")
	var produto models.Produto
	database.DB.First(&produto, id)
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&produto).UpdateColumns(produto)
	c.JSON(http.StatusOK, produto)
}

func DeletarProduto(c *gin.Context) {
	id := c.Params.ByName("id")
	var produto models.Produto
	database.DB.Delete(&produto, id)
	c.JSON(http.StatusOK, gin.H{"data": "Produto deletado com sucesso"})
}

func ComprarProduto(c *gin.Context) {
	var produtoParaComprar models.Produto
	if err := c.ShouldBindJSON(&produtoParaComprar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	var produto models.Produto
	database.DB.First(&produto, produtoParaComprar.GetId())
	if produto.GetQuantidade() > produtoParaComprar.GetQuantidade() {
		produto.SetQuantidade(produto.GetQuantidade() - produtoParaComprar.GetQuantidade())
		database.DB.Save(&produto)
		c.JSON(http.StatusOK, gin.H{"data": "Compra efetuada com sucesso"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Produto deletado com sucesso"})
	}
}
