package handlers

import (
	"errors"
	"log"
	base_handlers "symphony-api/internal/handlers/base"
	request_model "symphony-api/internal/handlers/model"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/repository"
	"symphony-api/internal/server"
	"symphony-api/internal/persistence/model"
)

type PostCrud struct {
	repository repository.PostRepository
	userRepository repository.UserRepository
}

func NewPostCrud(connection postgres.PostgreConnection) *PostCrud {
	return &PostCrud{
		userRepository: *repository.NewUserRepository(connection, nil),
		repository: *repository.NewPostRepository(connection),
	}
}

func (postCrud *PostCrud) AddRoutes(server server.Server) {
	server.AddRoute(
		"/api/post/create",
		base_handlers.CreatePostMethodHandler(postCrud.CreatePostHandler),
	)
	server.AddRoute(
		"/api/post/get-post-by-id",
		base_handlers.CreateGetMethodHandler(postCrud.GetPostByIdHandler),
	)
	server.AddRoute(
		"/api/post/get-by-username",
		base_handlers.CreateGetMethodHandler(postCrud.GetPostsByUsernameHandler),
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
//	@Router			/api/post/create [post]
func (postCrud *PostCrud) CreatePostHandler(request request_model.CreatePostRequest) (*request_model.CreatePostResponse, error) {
	log.Printf("CreatePostHandler called with request: %+v", request)
	user, err := postCrud.userRepository.GetByUsername(request.Username)

	if err != nil {
		log.Printf("Error creating post: %v", err)
		return nil, errors.New("error creating post")
	}

	createdPost, err := postCrud.repository.Put(
		&model.Post{
			UserId:    user.UserId,
			Text:      request.Text,
			UrlFoto:   request.UrlFoto,
			LikeCount: request.LikeCount,
		},
	)

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
//	@Router			/api/post/get-post-by-id [get]
func (postCrud *PostCrud) GetPostByIdHandler(request request_model.GetPostByIdRequest) (*request_model.GetPostByIdResponse, error) {
	post, err := postCrud.repository.GetById(request.PostId)
	if err != nil {
		log.Printf("Error getting post: %v", err)
		return nil, errors.New("error getting post")
	}
	return request_model.NewGetPostByIdResponse(post), nil
}

// GetPostsByUserIdHandler retrieves all posts for a specific user.
//	@Summary		Get posts by username
//	@Description	Retrieves all posts created by a specific user.
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			username	query		int	true	"Username"
//	@Success		200		{object}	request_model.GetPostsByUsernameResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/post/get-by-username [get]
func (postCrud *PostCrud) GetPostsByUsernameHandler(request request_model.GetPostsByUsernameRequest) (*request_model.GetPostsByUsernameResponse, error) {
	user, err := postCrud.userRepository.GetByUsername(request.Username)
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		return nil, errors.New("error getting posts")
	}
	posts, err := postCrud.repository.GetByUserId(user.UserId)
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		return nil, errors.New("error getting posts")
	}
	return request_model.NewGetPostsByUsernameResponse(posts), nil
}
