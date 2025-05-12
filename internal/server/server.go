package server

import (
	"log"
	"net/http"
	"symphony-api/internal/handlers"
)

// Run inicializa o servidor HTTP e mapeia as rotas para os handlers.
// A função mapeia a rota "/" para o handler HelloWorldHandler,
// que responde com "Hello, World!".
// A função retorna um erro se não for possível iniciar o servidor,
// caso contrário, o servidor continua em execução.
func Run() error {
	http.HandleFunc("/", handlers.HelloWorldHandler)

	log.Println("Server is running on http://localhost:8080")
	return http.ListenAndServe(":8080", nil)
}
