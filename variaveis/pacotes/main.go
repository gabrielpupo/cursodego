package main

import (
	"fmt"

	"github.com/gabrielpupo/cursodego/variaveis/pacotes/operadora"
	"github.com/gabrielpupo/cursodego/variaveis/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Jeff"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefiro da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
}
