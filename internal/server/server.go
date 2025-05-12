package server

import (
	"log"
	"net/http"
	"os"
	"symphony-api/internal/handlers"
)

// Run inicializa o servidor HTTP e mapeia as rotas para os handlers.
// A porta está na variável de ambiente "APP_PORT" ou padrão 8080.
// A função mapeia a rota "/" para o handler HelloWorldHandler,
// que responde com "Hello, World!".
// A função retorna um erro se não for possível iniciar o servidor,
// caso contrário, o servidor continua em execução.
func Run() error {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handlers.HelloWorldHandler)

	log.Printf("Server running on port %s\n", port)
	return http.ListenAndServe(":"+port, nil)
}
