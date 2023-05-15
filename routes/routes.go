package routes

import (
	"github.com/gin-gonic/gin"

	auth "loja/authentication"
	controllers "loja/controllers"
)

func HandleRequest() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") 
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	r.GET("/produtos", auth.TokenAuthMiddleware(), controllers.ListaTodosProdutos)
	r.GET("/produtos/:id", auth.TokenAuthMiddleware(), controllers.ObtemUmProduto)
	r.POST("/produtos", auth.TokenAuthMiddleware(), controllers.CriaNovoProduto)
	r.PUT("/produtos/:id", auth.TokenAuthMiddleware(), controllers.AtualizarProduto)
	r.DELETE("/produtos/:id", auth.TokenAuthMiddleware(), controllers.DeletarProduto)
	r.POST("/produtos/comprar", auth.TokenAuthMiddleware(), controllers.ComprarProduto)
	r.POST("/usuario/criar", controllers.CriaNovaConta)
	r.POST("/login", controllers.Login)
	r.Run(":8080")
}
