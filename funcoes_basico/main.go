package main

import (
	"fmt"

	"github.com/gabrielpupo/cursodego/funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
}

func multiplicador(x int, y int) int {
	return x * y
}
