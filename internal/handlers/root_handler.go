package handlers

import (
	"fmt"
	"net/http"
)

// RootHandler retorna um manipulador HTTP que responde com "Olá, mundo!".
// Este manipulador pode ser usado como a rota raiz do servidor.
func RootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá, mundo!")
	}
}
