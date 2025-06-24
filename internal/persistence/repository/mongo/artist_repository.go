package mongo_repository

import (
	"context"
	local_mongo "symphony-api/internal/persistence/connectors/mongo"

	"symphony-api/internal/persistence/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArtistRepository struct {
	collection *mongo.Collection
}

func NewArtistRepository(conn *local_mongo.MongoConnection) *ArtistRepository {
	coll := conn.GetCollection("symphony", "artists")
	return &ArtistRepository{collection: coll}
}

// InsertArtist insere um novo artista no banco.
func (r *ArtistRepository) InsertArtist(ctx context.Context, artist model.Artist) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, artist)
}

// GetArtistByID busca um artista pelo seu ID.
func (r *ArtistRepository) GetArtistByID(ctx context.Context, id primitive.ObjectID) (*model.Artist, error) {
	var artist model.Artist
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&artist)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

// GetArtistBySpotifyID busca um artista pelo seu ID no Spotify.
func (r *ArtistRepository) GetArtistBySpotifyID(ctx context.Context, idSpotify string) (*model.Artist, error) {
	var artist model.Artist
	err := r.collection.FindOne(ctx, bson.M{"id_spotify": idSpotify}).Decode(&artist)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}
