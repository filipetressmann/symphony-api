package server

import (
	"log"
	"net/http"
)

type Server struct {
	port string
	mux  *http.ServeMux
}

// NewServer creates a new instance of the Server struct.
// It initializes the server with the specified port and a new ServeMux.
// The ServeMux is used to route incoming HTTP requests to the appropriate handlers.
// The port parameter specifies the port on which the server will listen for incoming requests.
// The NewServer function returns a pointer to the newly created Server instance.
// It is designed to be used in a web application where you need to handle HTTP requests.
func NewServer(port string) *Server {
	return &Server{
		port: port,
		mux:  http.NewServeMux(),
	}
}

// AddRoute registers a new route with the server.
// It takes a path and a handler function as parameters.
// The path parameter specifies the URL path for the route,
// and the handler parameter is a function that will handle requests to that path.
// The handler function should conform to the http.HandlerFunc signature.
// This method allows you to define custom routes for your web application,
// enabling you to handle different HTTP methods and paths.
func (s *Server) AddRoute(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler)
}

// Start starts the HTTP server on the specified port.
// It listens for incoming HTTP requests and routes them to the appropriate handlers.
// The Start method blocks until the server is stopped or an error occurs.
// It uses the http.ListenAndServe function to start the server.
// If the server fails to start, it logs the error and exits the application.
// The port parameter specifies the port on which the server will listen for incoming requests.
func (s *Server) Start() {
	log.Printf("Starting server in port %s...", s.port)
	if err := http.ListenAndServe(":"+s.port, s.mux); err != nil {
		log.Fatalf("It wasn't possible to start the server in port %s\n", err)
	}
}
