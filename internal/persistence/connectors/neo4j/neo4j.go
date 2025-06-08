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

// InitNeo4j initializes a new Neo4j connection.
// It reads the connection parameters from environment variables,
// constructs the connection URI, and creates a Neo4j driver.
// If the connection fails, it logs the error and exits the application.
// The Neo4jConnection struct holds the Neo4j driver which can be used
// to interact with the Neo4j database.
// It also verifies the connectivity to the Neo4j database.
// If the connection is successful, it logs a success message.
// The connection parameters are:
// - NEO4J_HOST: The host of the Neo4j database (default: "neo4j").
// - NEO4J_PORT: The port of the Neo4j database (default: "7474").
// - NEO4J_USER: The username for the Neo4j database (default: "neo4j").
// - NEO4J_PASSWORD: The password for the Neo4j database (default: "password").
// Returns a pointer to a Neo4jConnection instance.
func NewNeo4jConnection() *Neo4jConnection {
	var client neo4j.Driver

	host := config.GetEnv("NEO4J_HOST", "neo4j")
	port := config.GetEnv("NEO4J_PORT", "7474")
	username := config.GetEnv("NEO4J_USER", "neo4j")
	password := config.GetEnv("NEO4J_PASSWORD", "password")

	uri := fmt.Sprintf("bolt://%s:%s", host, port)
	log.Println(uri)
	var err error
	client, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatalf("Failed to create neo4j driver: %v", err)
	}

	err = client.VerifyConnectivity()
	if err != nil {
		log.Fatalf("Failed to verify connectivity to neo4j: %v", err)
	}

	log.Println("Successfully connected to neo4j!")
	return &Neo4jConnection{
		client: &client,
	}
}
