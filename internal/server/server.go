package server

import (
	"log"
	"net/http"
)

type Server struct {
	port string
	mux  *http.ServeMux
}

// NewServer cria uma nova instância do servidor HTTP.
// O parâmetro 'port' especifica a porta na qual o servidor irá escutar.
func NewServer(port string) *Server {
	return &Server{
		port: port,
		mux:  http.NewServeMux(),
	}
}

// AddRoute adiciona uma nova rota ao servidor.
// O parâmetro 'path' especifica o caminho da rota e 'handler' é a função que irá lidar com as requisições para essa rota.
// O handler deve ser uma função que aceita um http.ResponseWriter e um http.Request como parâmetros.
func (s *Server) AddRoute(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler)
}

// Start inicia o servidor HTTP e escuta na porta especificada.
// Se houver um erro ao iniciar o servidor, ele será registrado e o programa será encerrado.
// O servidor irá escutar requisições na porta especificada e encaminhá-las para os manipuladores registrados.
func (s *Server) Start() {
	log.Printf("Iniciando servidor em %s...", s.port)
	if err := http.ListenAndServe(":"+s.port, s.mux); err != nil {
		log.Fatalf("Não foi possível iniciar o servidor: %s\n", err)
	}
}
