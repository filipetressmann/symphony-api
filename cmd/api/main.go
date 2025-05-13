package main

import (
	"log"
	"symphony-api/internal/server"
)


// main inicia a aplicação chamando o servidor HTTP.
// Ele configura e executa o servidor, e se ocorrer algum erro,
// o processo é interrompido e o erro é registrado.
func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}