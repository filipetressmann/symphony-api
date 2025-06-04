package handlers

import (
	"fmt"
	"net/http"
)

// RootHandler é um manipulador HTTP que responde com "Olá, mundo!".
// @Summary Manipulador raiz
// @Description Este manipulador é usado para verificar se o servidor está funcionando corretamente.
// @Tags Raiz
func RootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, "Olá, mundo!"); err != nil {
			http.Error(w, "Falha para escrever resposta.", http.StatusInternalServerError)
		}
	}
}
