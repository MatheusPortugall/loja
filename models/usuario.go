package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Id      uint64 `json:"id"`
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
}

func (u *Usuario) GetId() uint64 {
	return u.Id
}

func (u *Usuario) GetUsuario() string {
	return u.Usuario
}

func (u *Usuario) GetSenha() string {
	return u.Senha
}

func (u *Usuario) SetSenhaBcrypt(senha string) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(senha), 8)
	u.Senha = string(hashed)
}

func (u *Usuario) GetSenhaDecryptada() string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Senha), 8)
	// bcrypt.CompareHashAndPassword()
	return string(hashed)
}

var Usuarios []Usuario
