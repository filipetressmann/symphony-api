package main

import (
	"symphony-api/internal/handlers"
	chat_handlers "symphony-api/internal/handlers/chat"
	community_handlers "symphony-api/internal/handlers/community"
	user_handlers "symphony-api/internal/handlers/users"
	"symphony-api/internal/persistence/connectors/mongo"
	"symphony-api/internal/persistence/connectors/neo4j"

	//"symphony-api/internal/persistence/connectors/neo4j"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/server"
	"symphony-api/pkg/config"

	_ "symphony-api/tmp/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Symphony API
//	@version		1.0
//	@description	API for Symphony application, which is an social media created for educational purposes, focusing on music.
func main() {
	postgresConnection := postgres.NewPostgreConnection()
	_ = mongo.NewMongoConnection()
	neo4jConnection := neo4j.NewNeo4jConnection()
	
	userCrud := user_handlers.NewUserHandler(postgresConnection, neo4jConnection)
	postCrud := handlers.NewPostCrud(postgresConnection)
	communityCrud := community_handlers.NewCommunityHandler(postgresConnection, neo4jConnection)
	chatCrud := chat_handlers.NewChatHandler(postgresConnection, neo4jConnection)

	// Create a new server instance
	srv := server.NewServer(config.GetEnv("API_PORT", "8080"))

	srv.AddRoute("/", handlers.RootHandler())

	userCrud.AddRoutes(*srv)
	postCrud.AddRoutes(*srv)
	communityCrud.AddRoutes(*srv)
	chatCrud.AddRoutes(*srv)

	srv.AddRoute("/swagger/", httpSwagger.WrapHandler)

	srv.Start()
}
