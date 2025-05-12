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
	fmt.Fprintf(w, "Hello, World!")
}
