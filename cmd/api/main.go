package main

import (
	"symphony-api/internal/handlers"
	"symphony-api/internal/persistence/connectors/mongo"
	"symphony-api/internal/persistence/connectors/neo4j"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/server"
	"symphony-api/pkg/config"
)

// main é o ponto de entrada do aplicativo.
// Ele inicializa as conexões com os bancos de dados PostgreSQL, MongoDB e Neo4j,
// e cria uma nova instância do servidor HTTP.
// O servidor escuta na porta especificada pela variável de ambiente API_PORT ou na porta 8080 por padrão.
// Em seguida, adiciona rotas.
func main() {
	_ = postgres.NewPostgreConnection()
	_ = mongo.NewMongoConnection()
	_ = neo4j.NewNeo4jConnection()

	// Create a new server instance
	srv := server.NewServer(config.GetEnv("API_PORT", "8080"))

	srv.AddRoute("/", handlers.RootHandler())

	srv.Start()
}
