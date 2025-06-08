package main

import (
	"symphony-api/internal/handlers"
	community_handlers "symphony-api/internal/handlers/community"
	user_handlers "symphony-api/internal/handlers/users"
	"symphony-api/internal/persistence/connectors/mongo"

	//"symphony-api/internal/persistence/connectors/neo4j"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/server"
	"symphony-api/pkg/config"

	_ "symphony-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Symphony API
//	@version		1.0
//	@description	API for Symphony application, which is an social media created for educational purposes, focusing on music.
func main() {
	postgresConnection := postgres.NewPostgreConnection()
	_ = mongo.NewMongoConnection()
	//_ = neo4j.NewNeo4jConnection()

	userCrud := user_handlers.NewUserHandler(postgresConnection)
	postCrud := handlers.NewPostCrud(postgresConnection)
	communityCrud := community_handlers.NewCommunityHandler(postgresConnection)

	// Create a new server instance
	srv := server.NewServer(config.GetEnv("API_PORT", "8080"))

	srv.AddRoute("/", handlers.RootHandler())

	userCrud.AddRoutes(*srv)
	postCrud.AddRoutes(*srv)
	communityCrud.AddRoutes(*srv)

	srv.AddRoute("/swagger/", httpSwagger.WrapHandler)

	srv.Start()
}
