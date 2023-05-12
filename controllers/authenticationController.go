package controllers

import (
	"fmt"
	auth "loja/authentication"
	"loja/database"
	"loja/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var usuarioLogin models.Usuario
	if err := c.ShouldBindJSON(&usuarioLogin); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	var usuarioBanco models.Usuario
	database.DB.Where("usuario = ?", usuarioLogin.GetUsuario()).Find(&usuarioBanco)
	//b.Where("name <> ?", "jinzhu").Find(&users)
	fmt.Println("AQUI: " + usuarioBanco.GetUsuario())

	if usuarioBanco.GetUsuario() != usuarioLogin.GetUsuario() || usuarioBanco.GetSenha() != usuarioLogin.GetSenha() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": true, "message": "Nâo autorizado"})
		return
	}
	token, err := auth.CreateToken(usuarioLogin.GetId())
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func CriaNovaConta(c *gin.Context) {
	var u models.Usuario
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	senhaSemHash := u.GetSenha()
	u.SetSenhaBcrypt(senhaSemHash)
	database.DB.Create(&u)
	ret := bcrypt.CompareHashAndPassword([]byte(u.GetSenha()), []byte(senhaSemHash))
	fmt.Println(ret)
	c.JSON(http.StatusOK, u)
}
