package playlist

import (
	"context"
	"encoding/json"
	"net/http"
	"symphony-api/internal/persistence/model"
	mongo_repository "symphony-api/internal/persistence/repository/mongo"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlaylistHandler struct {
	repo *mongo_repository.PlaylistRepository
}

func NewPlaylistHandler(repo *mongo_repository.PlaylistRepository) *PlaylistHandler {
	return &PlaylistHandler{repo: repo}
}

func (h *PlaylistHandler) AddRoutes(server interface {
	AddRoute(pattern string, handler http.HandlerFunc)
	AddGroup(pattern string, fn func(r chi.Router))
}) {
	server.AddGroup("/playlists", func(r chi.Router) {
		r.Get("/{id}", h.GetPlaylistByID)
		r.Get("/user/{username}", h.GetPlaylistsByUsername)
		r.Post("/", h.CreatePlaylist)
	})
}

// GetPlaylistByID returns a playlist by its ID
// @Summary Get playlist by ID
// @Description Get a playlist by its MongoDB ObjectID
// @Tags playlists
// @Accept json
// @Produce json
// @Param id path string true "Playlist ID"
// @Success 200 {object} model.Playlist
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /playlists/{id} [get]
func (h *PlaylistHandler) GetPlaylistByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	idStr := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid playlist ID", http.StatusBadRequest)
		return
	}

	playlist, err := h.repo.GetPlaylistByID(ctx, id)
	if err != nil {
		http.Error(w, "Playlist not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(playlist)
}

// GetPlaylistsByUsername returns all playlists created by a user
// @Summary Get user's playlists
// @Description Get all playlists created by a specific user
// @Tags playlists
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {array} model.Playlist
// @Failure 404 {object} map[string]string
// @Router /playlists/user/{username} [get]
func (h *PlaylistHandler) GetPlaylistsByUsername(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	username := chi.URLParam(r, "username")

	playlists, err := h.repo.GetPlaylistsByUserID(ctx, username)
	if err != nil {
		http.Error(w, "Playlists not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(playlists)
}

// CreatePlaylistRequest represents the request body for creating a new playlist
type CreatePlaylistRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	Public      bool   `json:"public,omitempty"`
	IDSpotify   string `json:"id_spotify,omitempty"`
	Title       string `json:"title,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	Songs       []struct {
		SongID string `json:"song_id,omitempty"`
		Order  int32  `json:"order,omitempty"`
	} `json:"songs,omitempty"`
}

// CreatePlaylist creates a new playlist
// @Summary Create a new playlist
// @Description Create a new playlist in the database
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist body CreatePlaylistRequest true "Playlist object"
// @Success 201 {object} model.Playlist
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /playlists [post]
func (h *PlaylistHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req CreatePlaylistRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Convert string song IDs to ObjectIDs if provided
	songs := make([]struct {
		SongID primitive.ObjectID `bson:"song_id,omitempty"`
		Order  int32              `bson:"order,omitempty"`
	}, 0, len(req.Songs))

	for _, song := range req.Songs {
		songID, err := primitive.ObjectIDFromHex(song.SongID)
		if err != nil {
			http.Error(w, "Invalid song ID", http.StatusBadRequest)
			return
		}
		songs = append(songs, struct {
			SongID primitive.ObjectID `bson:"song_id,omitempty"`
			Order  int32              `bson:"order,omitempty"`
		}{
			SongID: songID,
			Order:  song.Order,
		})
	}

	now := time.Now()
	playlist := model.Playlist{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		UserID:      req.UserID,
		Public:      req.Public,
		IDSpotify:   req.IDSpotify,
		Title:       req.Title,
		ImageURL:    req.ImageURL,
		CreatedAt:   now,
		UpdatedAt:   now,
		Songs:       songs,
	}

	_, err := h.repo.InsertPlaylist(ctx, playlist)
	if err != nil {
		http.Error(w, "Failed to insert playlist", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(playlist)
}
