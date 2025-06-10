package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	base_handlers "symphony-api/internal/handlers/base"
	request_model "symphony-api/internal/handlers/model"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/repository"
	"symphony-api/internal/server"
)

type PostCrud struct {
	repository repository.PostRepository
}

func NewPostCrud(connection postgres.PostgreConnection) *PostCrud {
	return &PostCrud{
		repository: *repository.NewPostRepository(connection),
	}
}

func (postCrud *PostCrud) AddRoutes(server server.Server) {
	server.AddRoute(
		"/api/create-post",
		base_handlers.CreateHandler(postCrud.CreatePostHandler),
	)
	server.AddRoute(
		"/api/get-post-by-id",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			postId, err := strconv.ParseInt(r.URL.Query().Get("post_id"), 10, 32)
			if err != nil {
				http.Error(w, "Invalid post ID", http.StatusBadRequest)
				return
			}
			request := request_model.GetPostByIdRequest{PostId: int32(postId)}
			response, err := postCrud.GetPostByIdHandler(request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Error encoding response", http.StatusInternalServerError)
				return
			}
		}),
	)
	server.AddRoute(
		"/api/get-posts-by-user-id",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 32)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}
			request := request_model.GetPostsByUserIdRequest{UserId: int32(userId)}
			response, err := postCrud.GetPostsByUserIdHandler(request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Error encoding response", http.StatusInternalServerError)
				return
			}
		}),
	)
}

// CreatePostHandler handles the creation of a new post.
//	@Summary		Create a new post
//	@Description	Creates a new post in the system.
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			post	body		request_model.CreatePostRequest	true	"Post data"
//	@Success		200		{object}	request_model.CreatePostResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/create-post [post]
func (postCrud *PostCrud) CreatePostHandler(request request_model.CreatePostRequest) (*request_model.CreatePostResponse, error) {
	log.Printf("CreatePostHandler called with request: %+v", request)
	createdPost, err := postCrud.repository.Put(request.ToPost())

	if err != nil {
		log.Printf("Error creating post: %v", err)
		return nil, errors.New("error creating post")
	}

	return request_model.NewCreatePostResponse(createdPost), nil
}

// GetPostByIdHandler retrieves a post by its ID.
//	@Summary		Get post by ID
//	@Description	Retrieves a post using its unique identifier.
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			post_id	query		int	true	"Post ID"
//	@Success		200		{object}	request_model.GetPostByIdResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		404		{object}	map[string]string	"Post Not Found"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/get-post-by-id [get]
func (postCrud *PostCrud) GetPostByIdHandler(request request_model.GetPostByIdRequest) (*request_model.GetPostByIdResponse, error) {
	post, err := postCrud.repository.GetById(request.PostId)
	if err != nil {
		log.Printf("Error getting post: %v", err)
		return nil, errors.New("error getting post")
	}
	return request_model.NewGetPostByIdResponse(post), nil
}

// GetPostsByUserIdHandler retrieves all posts for a specific user.
//	@Summary		Get posts by user ID
//	@Description	Retrieves all posts created by a specific user.
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		int	true	"User ID"
//	@Success		200		{object}	request_model.GetPostsByUserIdResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/get-posts-by-user-id [get]
func (postCrud *PostCrud) GetPostsByUserIdHandler(request request_model.GetPostsByUserIdRequest) (*request_model.GetPostsByUserIdResponse, error) {
	posts, err := postCrud.repository.GetByUserId(request.UserId)
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		return nil, errors.New("error getting posts")
	}
	return request_model.NewGetPostsByUserIdResponse(posts), nil
}
