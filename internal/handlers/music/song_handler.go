package music

import (
	"context"
	"encoding/json"
	"net/http"
	"symphony-api/internal/persistence/model"
	mongo_repository "symphony-api/internal/persistence/repository/mongo"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SongHandler struct {
	repo *mongo_repository.SongRepository
}

func NewSongHandler(repo *mongo_repository.SongRepository) *SongHandler {
	return &SongHandler{repo: repo}
}

func (h *SongHandler) AddRoutes(server interface {
	AddRoute(pattern string, handler http.HandlerFunc)
	AddGroup(pattern string, fn func(r chi.Router))
}) {
	server.AddGroup("/songs", func(r chi.Router) {
		r.Get("/", h.GetAllSongs)
		r.Get("/{id}", h.GetSongByID)
		r.Post("/", h.CreateSong)
	})
}

// GetAllSongs returns all songs in the database
// @Summary Get all songs
// @Description Get a list of all songs
// @Tags songs
// @Accept json
// @Produce json
// @Success 200 {array} model.Song
// @Failure 500 {object} map[string]string
// @Router /songs [get]
func (h *SongHandler) GetAllSongs(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	songs, err := h.repo.GetAllSongs(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch songs", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(songs); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetSongByID returns a song by its ID
// @Summary Get a song by ID
// @Description Get a song by its MongoDB ObjectID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Success 200 {object} model.Song
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /songs/{id} [get]
func (h *SongHandler) GetSongByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	idStr := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	song, err := h.repo.GetSongByID(ctx, id)
	if err != nil {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(song); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// CreateSongRequest represents the request body for creating a new song
type CreateSongRequest struct {
	IDSpotify   string `json:"id_spotify,omitempty"`
	Title       string `json:"title,omitempty"`
	Duration    int32  `json:"duration,omitempty"`
	ArtistID    string `json:"artist_id,omitempty"`
	Genre       string `json:"genre,omitempty"`
	ReseaseYear int32  `json:"release_year,omitempty"`
	Album       string `json:"album,omitempty"`
	URLSpotify  string `json:"url_spotify,omitempty"`
}

// CreateSong creates a new song
// @Summary Create a new song
// @Description Create a new song in the database
// @Tags songs
// @Accept json
// @Produce json
// @Param song body CreateSongRequest true "Song object"
// @Success 201 {object} model.Song
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /songs [post]
func (h *SongHandler) CreateSong(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req CreateSongRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Convert string artist ID to ObjectID if provided
	var artistID primitive.ObjectID
	if req.ArtistID != "" {
		var err error
		artistID, err = primitive.ObjectIDFromHex(req.ArtistID)
		if err != nil {
			http.Error(w, "Invalid artist ID", http.StatusBadRequest)
			return
		}
	}

	song := model.Song{
		ID:          primitive.NewObjectID(),
		IDSpotify:   req.IDSpotify,
		Title:       req.Title,
		Duration:    req.Duration,
		ArtistID:    artistID,
		Genre:       req.Genre,
		ReseaseYear: req.ReseaseYear,
		Album:       req.Album,
		URLSpotify:  req.URLSpotify,
	}

	_, err := h.repo.InsertSong(ctx, song)
	if err != nil {
		http.Error(w, "Failed to insert song", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(song); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
