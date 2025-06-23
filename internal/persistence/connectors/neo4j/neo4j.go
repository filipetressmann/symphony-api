package neo4j

import (
	"context"
	"fmt"
	"symphony-api/pkg/config"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jConnection interface {
	Execute(query string, data map[string]any) (error)
	ExecuteReturning(query string, data map[string]any) ([]*neo4j.Record, error)
}

type Neo4jConnectionImpl struct {
	client neo4j.DriverWithContext
	ctx context.Context
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
// - NEO4J_PORT: The port of the Neo4j database (default: "7687").
// - NEO4J_USER: The username for the Neo4j database (default: "neo4j").
// - NEO4J_PASSWORD: The password for the Neo4j database (default: "password").
// Returns a pointer to a Neo4jConnection instance.
func NewNeo4jConnection() Neo4jConnection {
	var client neo4j.DriverWithContext

	ctx := context.Background()
	
	uri := config.GetEnv("NEO4J_HOST", "neo4j")
    username := config.GetEnv("NEO4J_USER", "neo4j")
	password := config.GetEnv("NEO4J_PASSWORD", "neo4j")
    client, err := neo4j.NewDriverWithContext(
        uri,
        neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic(err)
	}
    
    err = client.VerifyConnectivity(ctx)
    if err != nil {
        panic(err)
    }
    fmt.Println("Connection established to neo4j.")

	return &Neo4jConnectionImpl{
		client: client,
		ctx: ctx,
	}
}

func (connection *Neo4jConnectionImpl) Execute(query string, data map[string]any) (error) {
	_, err := neo4j.ExecuteQuery(
		connection.ctx, 
		connection.client, 
		query, 
		data, 
		neo4j.EagerResultTransformer, 
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)

	return err
}

func (connection *Neo4jConnectionImpl) ExecuteReturning(query string, data map[string]any) ([]*neo4j.Record, error) {
	result, err := neo4j.ExecuteQuery(
		connection.ctx, 
		connection.client, 
		query, 
		data, 
		neo4j.EagerResultTransformer, 
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	return result.Records, err
}