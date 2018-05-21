package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Qual o valor de i?", i)
	}

	valor := 0
	teste := true

	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("O numero é: ", valor)
	}

}
