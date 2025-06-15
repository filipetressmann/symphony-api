package main

import (
	"symphony-api/internal/handlers"
	community_handlers "symphony-api/internal/handlers/community"
	music_handlers "symphony-api/internal/handlers/music"
	artist_handlers "symphony-api/internal/handlers/artist"
	user_handlers "symphony-api/internal/handlers/users"
	"symphony-api/internal/persistence/connectors/mongo"
	"symphony-api/internal/persistence/connectors/postgres"
	mongo_repository "symphony-api/internal/persistence/repository/mongo"
	"symphony-api/internal/server"
	"symphony-api/pkg/config"

	_ "symphony-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	postgresConnection := postgres.NewPostgreConnection()
	mongoConnection := mongo.NewMongoConnection()

	// Repositórios
	songRepo := mongo_repository.NewSongRepository(mongoConnection)
	artistRepo := mongo_repository.NewArtistRepository(mongoConnection)

	// Handlers
	userCrud := user_handlers.NewUserHandler(postgresConnection)
	postCrud := handlers.NewPostCrud(postgresConnection)
	communityCrud := community_handlers.NewCommunityHandler(postgresConnection)
	songHandler := music_handlers.NewSongHandler(songRepo)
	artistHandler := artist_handlers.NewArtistHandler(artistRepo)

	// Servidor
	srv := server.NewServer(config.GetEnv("API_PORT", "8080"))

	srv.AddRoute("/", handlers.RootHandler())

	userCrud.AddRoutes(*srv)
	postCrud.AddRoutes(*srv)
	communityCrud.AddRoutes(*srv)
	songHandler.AddRoutes(srv)
	artistHandler.AddRoutes(srv)

	// Swagger
	srv.AddRoute("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	srv.Start()
}
