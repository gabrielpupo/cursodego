package main

import (
	"fmt"
)

func main() {

	meses := 0
	situacao := true
	cidade := "Teste"

	//< > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo.")
	}

	if situacao {
		fmt.Println("Ele está devendo")
	}

	if !situacao {
		fmt.Println("Ele está adimplente")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive no estado Teste")
	}

	if descricao, status := haQuantoTenmpoEstaDevendo(meses); status {
		fmt.Println("Qual a situacao do cliente?", descricao)
		return
	}

	fmt.Println("Obrigado por nos consultar")
}

func haQuantoTenmpoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo"
		return
	}

	descricao = "O cliente está em dia"
	return
}
