package main

import (
	"fmt"
)

var (
	//Endereco é um valor importante e tem de ser publico
	Endereco string
	//Telefone é um valor importante para a nossa aula
	telefone            string
	quantidade, estoque int
	comprou             bool
	valor               float64
	palavras            rune
)

func main() {
	teste := "valor de teste"
	fmt.Printf("Endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("O valor de teste é: %v\r\n", teste)
}
