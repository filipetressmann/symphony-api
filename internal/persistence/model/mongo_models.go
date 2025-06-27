package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	ID          primitive.ObjectID 		`bson:"_id"`
	IDSpotify   string             		`bson:"id_spotify,omitempty"`
	Title       string             		`bson:"title,omitempty"`
	Duration    int32              		`bson:"duration,omitempty"`
	ArtistID    primitive.ObjectID 		`bson:"artist_id,omitempty"`
	Genre       string             		`bson:"genre,omitempty"`
	ReseaseYear int32          			`bson:"release_year,omitempty"`
	Album       string         			`bson:"album,omitempty"`
	URLSpotify  string         			`bson:"url_spotify,omitempty"`
}

type Artist struct {
	ID          primitive.ObjectID `bson:"_id"`
	IDSpotify   string             `bson:"id_spotify,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Country     string             `bson:"country,omitempty"`
	ImageURL    string             `bson:"image_url,omitempty"`
	Biography   string         		`bson:"biography,omitempty"`
	Genres      []string       		`bson:"genres,omitempty"`
}

type Playlist struct {
	ID          primitive.ObjectID 	`bson:"_id"`
	Name        string             	`bson:"name,omitempty"`
	Public      bool               	`bson:"public,omitempty"`
	Username    string             	`bson:"username,omitempty"`
	IDSpotify   string             	`bson:"id_spotify,omitempty"`
	Title       string             	`bson:"title,omitempty"`
	Description string             	`bson:"description,omitempty"`
	ImageURL    string             	`bson:"image_url,omitempty"`
	CreatedAt   time.Time      		`bson:"created_at,omitempty"`
	UpdatedAt   time.Time      		`bson:"updated_at,omitempty"`
	Likes       []string       		`bson:"likes,omitempty"`
	Songs       []struct {
		SongID primitive.ObjectID 		`bson:"song_id,omitempty"`
		Order  int32              		`bson:"order,omitempty"`
	} `bson:"songs,omitempty"`
}
