package neo4j

import (
	"fmt"
	"log"
	"symphony-api/pkg/config"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jConnection struct {
	client *neo4j.Driver
}

// InitNeo4j inicializa a conexão com o banco de dados Neo4j e retorna o cliente.
// O cliente pode ser usado para executar consultas e interagir com o banco de dados.
// O URL de conexão é construído a partir das variáveis de ambiente definidas.
// As variáveis de ambiente esperadas são:
// NEO4J_HOST: Endereço do host do Neo4j (padrão: "localhost")
// NEO4J_PORT: Porta do Neo4j (padrão: "7687")
// NEO4J_USER: Nome de usuário do Neo4j (padrão: "neo4j")
// NEO4J_PASSWORD: Senha do Neo4j (padrão: "password")
// Se a conexão falhar, o programa será encerrado com um log de erro.
func NewNeo4jConnection() *Neo4jConnection {
	var client neo4j.Driver

	host := config.GetEnv("NEO4J_HOST", "localhost")
	port := config.GetEnv("NEO4J_PORT", "7687")
	username := config.GetEnv("NEO4J_USER", "neo4j")
	password := config.GetEnv("NEO4J_PASSWORD", "password")

	uri := fmt.Sprintf("bolt://%s:%s", host, port)

	var err error
	client, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatalf("Falha ao criar driver do Neo4j: %v", err)
	}

	err = client.VerifyConnectivity()
	if err != nil {
		log.Fatalf("Falha ao verificar Neo4j: %v", err)
	}
	return &Neo4jConnection{
		client: &client,
	}
}
