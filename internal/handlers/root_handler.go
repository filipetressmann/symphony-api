package handlers

import (
	"fmt"
	"net/http"
)

// RootHandler retorna um manipulador HTTP que responde com "Olá, mundo!".
// Este manipulador pode ser usado como a rota raiz do servidor.
func RootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, "Olá, mundo!"); err != nil {
			http.Error(w, "Falha para escrever resposta.", http.StatusInternalServerError)
		}
	}
}
