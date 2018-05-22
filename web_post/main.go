package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gabrielpupo/cursodego/web_post/model"
)

func main() {

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Jeff Prestes"

	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao gerar o JSON do objeto usuario", err.Error())
		return
	}

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}
	request, err := http.NewRequest("POST", "http://requestbin.fullcontact.com/ul2rxbul", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request da pagina do Request Bin", err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao chamar o servico do requestbin", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo retornado pelo request bin", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}
}
