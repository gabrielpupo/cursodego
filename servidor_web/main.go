package main

import (
	"fmt"
	"net/http"

	"github.com/gabrielpupo/cursodego/servidor_web/manipulador"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá Mundo")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("Servidor rodando na porta 8181")
	http.ListenAndServe(":8181", nil)

}
