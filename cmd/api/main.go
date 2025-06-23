package main

import (
	"symphony-api/internal/handlers"
	chat_handlers "symphony-api/internal/handlers/chat"
	community_handlers "symphony-api/internal/handlers/community"
	user_handlers "symphony-api/internal/handlers/users"
	"symphony-api/internal/persistence/connectors/mongo"
	music_handlers "symphony-api/internal/handlers/music"
	playlist_handlers "symphony-api/internal/handlers/playlist"
	artist_handlers "symphony-api/internal/handlers/artist"
	mongo_repository "symphony-api/internal/persistence/repository/mongo"
	"symphony-api/internal/persistence/connectors/neo4j"

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
	mongoConnection := mongo.NewMongoConnection()
  neo4jConnection := neo4j.NewNeo4jConnection()

	// Repositórios
	songRepo := mongo_repository.NewSongRepository(mongoConnection)
	artistRepo := mongo_repository.NewArtistRepository(mongoConnection)
	playlistRepo := mongo_repository.NewPlaylistRepository(mongoConnection)

	// Handlers
	userCrud := user_handlers.NewUserHandler(postgresConnection, neo4jConnection)
	postCrud := handlers.NewPostCrud(postgresConnection)
	communityCrud := community_handlers.NewCommunityHandler(postgresConnection, neo4jConnection)
  chatCrud := chat_handlers.NewChatHandler(postgresConnection, neo4jConnection)
  songHandler := music_handlers.NewSongHandler(songRepo)
	artistHandler := artist_handlers.NewArtistHandler(artistRepo)
	playlistHandler := playlist_handlers.NewPlaylistHandler(playlistRepo)

	// Create a new server instance
	srv := server.NewServer(config.GetEnv("API_PORT", "8080"))

	srv.AddRoute("/", handlers.RootHandler())

	userCrud.AddRoutes(*srv)
	postCrud.AddRoutes(*srv)
	communityCrud.AddRoutes(*srv)
	chatCrud.AddRoutes(*srv)
	songHandler.AddRoutes(srv)
	artistHandler.AddRoutes(srv)
	playlistHandler.AddRoutes(srv)

	// Swagger
	srv.AddRoute("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	srv.Start()
}
