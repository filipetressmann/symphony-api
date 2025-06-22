package artist

import (
	"context"
	"encoding/json"
	"net/http"
	"symphony-api/internal/persistence/model"
	mongo_repository "symphony-api/internal/persistence/repository/mongo"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArtistHandler struct {
	repo *mongo_repository.ArtistRepository
}

func NewArtistHandler(repo *mongo_repository.ArtistRepository) *ArtistHandler {
	return &ArtistHandler{repo: repo}
}

func (h *ArtistHandler) AddRoutes(server interface {
	AddRoute(pattern string, handler http.HandlerFunc)
	AddGroup(pattern string, fn func(r chi.Router))
}) {
	server.AddGroup("/artists", func(r chi.Router) {
		r.Get("/{id}", h.GetArtistByID)
		r.Get("/spotify/{spotify_id}", h.GetArtistBySpotifyID)
		r.Post("/", h.CreateArtist)
	})
}

// GetArtistByID returns an artist by its ObjectID
// @Summary Get artist by ID
// @Description Get artist data by MongoDB ObjectID
// @Tags artists
// @Accept json
// @Produce json
// @Param id path string true "Artist ObjectID"
// @Success 200 {object} model.Artist
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /artists/{id} [get]
func (h *ArtistHandler) GetArtistByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	idStr := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artist, err := h.repo.GetArtistByID(ctx, id)
	if err != nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(artist); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetArtistBySpotifyID returns an artist by their Spotify ID
// @Summary Get artist by Spotify ID
// @Description Get artist data by Spotify ID
// @Tags artists
// @Accept json
// @Produce json
// @Param spotify_id path string true "Spotify Artist ID"
// @Success 200 {object} model.Artist
// @Failure 404 {object} map[string]string
// @Router /artists/spotify/{spotify_id} [get]
func (h *ArtistHandler) GetArtistBySpotifyID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	idSpotify := chi.URLParam(r, "spotify_id")

	artist, err := h.repo.GetArtistBySpotifyID(ctx, idSpotify)
	if err != nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(artist); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// CreateArtistRequest represents the request body for creating a new artist
type CreateArtistRequest struct {
	IDSpotify   string   `json:"id_spotify,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Country     string   `json:"country,omitempty"`
	ImageURL    string   `json:"image_url,omitempty"`
	Biography   string   `json:"biography,omitempty"`
	Genres      []string `json:"genres,omitempty"`
}

// CreateArtist creates a new artist
// @Summary Create a new artist
// @Description Create a new artist in the database
// @Tags artists
// @Accept json
// @Produce json
// @Param artist body CreateArtistRequest true "Artist object"
// @Success 201 {object} model.Artist
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /artists [post]
func (h *ArtistHandler) CreateArtist(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req CreateArtistRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	artist := model.Artist{
		ID:          primitive.NewObjectID(),
		IDSpotify:   req.IDSpotify,
		Name:        req.Name,
		Description: req.Description,
		Country:     req.Country,
		ImageURL:    req.ImageURL,
		Biography:   req.Biography,
		Genres:      req.Genres,
	}

	_, err := h.repo.InsertArtist(ctx, artist)
	if err != nil {
		http.Error(w, "Failed to create artist", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(artist); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
