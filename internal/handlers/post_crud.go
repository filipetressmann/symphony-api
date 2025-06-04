package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/model"
	"symphony-api/internal/persistence/repository"
)

type PostCrud struct {
	repository repository.PostRepository
}

func NewPostCrud(connection postgres.PostgreConnection) *PostCrud {
	return &PostCrud{
		repository: *repository.NewPostRepository(connection),
	}
}

func (postCrud *PostCrud) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	post, err := model.PostFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	createdPost, err := postCrud.repository.Put(post)
	if err != nil {
		log.Printf("Error creating post: %s", err)
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Post created successfully",
		"post":    createdPost,
	})

	if err != nil {
		log.Printf("Error processing answer: %s", err)
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}
}

func (postCrud *PostCrud) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.ParseInt(r.URL.Query().Get("post_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post := postCrud.repository.GetById(postId)
	if post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		log.Printf("Error processing answer: %s", err)
		http.Error(w, "Error retrieving post", http.StatusInternalServerError)
		return
	}
}

func (postCrud *PostCrud) GetUserPostsHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	posts := postCrud.repository.GetByUserId(userId)
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Printf("Error processing answer: %s", err)
		http.Error(w, "Error retrieving posts", http.StatusInternalServerError)
		return
	}
}
