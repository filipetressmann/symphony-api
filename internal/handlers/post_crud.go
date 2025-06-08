package handlers

import (
	"errors"
	"log"
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
		createHandler(postCrud.CreatePostHandler),
	)
	server.AddRoute(
		"/api/get-post-by-id",
		createHandler(postCrud.GetPostByIdHandler),
	)
	server.AddRoute(
		"/api/get-posts-by-user-id",
		createHandler(postCrud.GetPostsByUserIdHandler),
	)
}

func (postCrud *PostCrud) CreatePostHandler(request request_model.CreatePostRequest) (*request_model.CreatePostResponse, error) {
	log.Printf("CreatePostHandler called with request: %+v", request)
	createdPost, err := postCrud.repository.Put(request.ToPost())

	if err != nil {
		log.Printf("Error creating post: %v", err)
		return nil, errors.New("error creating post")
	}

	return request_model.NewCreatePostResponse(createdPost), nil
}

func (postCrud *PostCrud) GetPostByIdHandler(request request_model.GetPostByIdRequest) (*request_model.GetPostByIdResponse, error) {
	post, err := postCrud.repository.GetById(request.PostId)
	if err != nil {
		log.Printf("Error getting post: %v", err)
		return nil, errors.New("error getting post")
	}
	return request_model.NewGetPostByIdResponse(post), nil
}

func (postCrud *PostCrud) GetPostsByUserIdHandler(request request_model.GetPostsByUserIdRequest) (*request_model.GetPostsByUserIdResponse, error) {
	posts, err := postCrud.repository.GetByUserId(request.UserId)
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		return nil, errors.New("error getting posts")
	}
	return request_model.NewGetPostsByUserIdResponse(posts), nil
}
