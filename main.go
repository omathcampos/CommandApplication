package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}

}

// Para rodar no terminal entre na pasta onde está a aplicação e para buscar por servidores
//use - go run main.go Servidores --host url
// e para buscar por IP use go run main.go IP --host nome do site

// Com o go build feito use ./linha-de-comando Servidores (ou IP) host-- url
