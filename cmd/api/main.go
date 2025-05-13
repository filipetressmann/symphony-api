package main

import (
	"symphony-api/internal/handlers"
	"symphony-api/internal/mongo"
	"symphony-api/internal/neo4j"
	"symphony-api/internal/postgres"
	"symphony-api/internal/server"
	"symphony-api/pkg/config"
)

// main é o ponto de entrada do aplicativo.
// Ele inicializa as conexões com os bancos de dados PostgreSQL, MongoDB e Neo4j,
// e cria uma nova instância do servidor HTTP.
// O servidor escuta na porta especificada pela variável de ambiente API_PORT ou na porta 8080 por padrão.
// Em seguida, adiciona rotas.
func main() {
	// Initialize database connections
	_ = postgres.InitPostgres() // pgConn
	_ = mongo.InitMongo()       // mongoConn
	_ = neo4j.InitNeo4j()       // neo4jConn

	// Create a new server instance
	srv := server.NewServer(config.GetEnv("API_PORT", "8080"))

	srv.AddRoute("/", handlers.RootHandler())

	srv.Start()
}
