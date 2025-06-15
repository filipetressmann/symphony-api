// internal/persistence/repository/mongo/song_repository.go
package mongo_repository

import (
	"context"
	local_mongo "symphony-api/internal/persistence/connectors/mongo"
	"symphony-api/internal/persistence/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SongRepository struct {
	collection *mongo.Collection
}

func NewSongRepository(conn *local_mongo.MongoConnection) *SongRepository {
	coll := conn.GetCollection("symphony", "songs")
	return &SongRepository{collection: coll}
}

// InsertSong insere uma nova música no banco.
func (r *SongRepository) InsertSong(ctx context.Context, song model.Song) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, song)
}

// GetSongByID busca uma música pelo seu ID.
func (r *SongRepository) GetSongByID(ctx context.Context, id primitive.ObjectID) (*model.Song, error) {
	var song model.Song
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&song)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

// GetAllSongs lista todas as músicas.
func (r *SongRepository) GetAllSongs(ctx context.Context) ([]model.Song, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var songs []model.Song
	for cursor.Next(ctx) {
		var song model.Song
		if err := cursor.Decode(&song); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return songs, nil
}
