package handlers

import (
	"fmt"
	"net/http"
)

// HelloWorldHandler é o handler para a rota "/".
// Ele escreve "Hello, World!" no corpo da resposta HTTP.
// Essa função não recebe parâmetros e não retorna nenhum valor.
// Ela é responsável por responder ao cliente com a mensagem simples.
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
