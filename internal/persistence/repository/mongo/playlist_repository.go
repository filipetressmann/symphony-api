package mongo_repository

import (
	"context"
	local_mongo "symphony-api/internal/persistence/connectors/mongo"

	"symphony-api/internal/persistence/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlaylistRepository struct {
	collection *mongo.Collection
}

func NewPlaylistRepository(conn *local_mongo.MongoConnection) *PlaylistRepository {
	coll := conn.GetCollection("symphony", "playlists")
	return &PlaylistRepository{collection: coll}
}

func (r *PlaylistRepository) InsertPlaylist(ctx context.Context, playlist model.Playlist) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, playlist)
}

func (r *PlaylistRepository) GetPlaylistByID(ctx context.Context, id primitive.ObjectID) (*model.Playlist, error) {
	var playlist model.Playlist
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&playlist)
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (r *PlaylistRepository) GetPlaylistsByUserID(ctx context.Context, userID string) ([]model.Playlist, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := cursor.Close(ctx); closeErr != nil {
			err = closeErr
		}
	}()

	var playlists []model.Playlist
	for cursor.Next(ctx) {
		var playlist model.Playlist
		if err := cursor.Decode(&playlist); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return playlists, nil
}

// UpdatePlaylist updates an existing playlist in the database
func (r *PlaylistRepository) UpdatePlaylist(ctx context.Context, id primitive.ObjectID, playlist model.Playlist) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": id}, playlist)
	return err
}
