package model

//Usuario Representa o usuario de acesso a aplicacao
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}
