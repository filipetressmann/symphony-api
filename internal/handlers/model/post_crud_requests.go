package request_model

import (
	"symphony-api/internal/persistence/model"
)

type CreatePostRequest struct {
	UserId int32 `json:"user_id" binding:"required"`
	*BasePostModel
}

func (request *CreatePostRequest) ToPost() *model.Post {
	return &model.Post{
		UserId:    request.UserId,
		Text:      request.Text,
		UrlFoto:   request.UrlFoto,
		LikeCount: request.LikeCount,
	}
}

type BasePostModel struct {
	Text      string `json:"text" binding:"required"`
	UrlFoto   string `json:"url_foto" binding:"required"`
	LikeCount int    `json:"like_count" binding:"required"`
}

func NewBasePostModel(post *model.Post) *BasePostModel {
	return &BasePostModel{
		Text:      post.Text,
		UrlFoto:   post.UrlFoto,
		LikeCount: post.LikeCount,
	}
}

type CreatePostResponse struct {
	*BasePostModel
}

func (request *CreatePostResponse) ToPost() *model.Post {
	return &model.Post{
		Text:      request.Text,
		UrlFoto:   request.UrlFoto,
		LikeCount: request.LikeCount,
	}
}

func NewCreatePostResponse(post *model.Post) *CreatePostResponse {
	return &CreatePostResponse{
		BasePostModel: NewBasePostModel(post),
	}
}

type PostResponse struct {
	*BasePostModel
	Id int32 `json:"id" binding:"required"`
}

func NewPostResponse(post *model.Post) *PostResponse {
	return &PostResponse{
		Id:            post.PostId,
		BasePostModel: NewBasePostModel(post),
	}
}

type GetPostByIdRequest struct {
	PostId int32 `json:"post_id" binding:"required"`
}

type GetPostByIdResponse struct {
	Id     int32 `json:"id" binding:"required"`
	UserId int32 `json:"user_id" binding:"required"`
	*BasePostModel
}

func NewGetPostByIdResponse(post *model.Post) *GetPostByIdResponse {
	return &GetPostByIdResponse{
		Id:            post.PostId,
		UserId:        post.UserId,
		BasePostModel: NewBasePostModel(post),
	}
}

type GetPostsByUserIdRequest struct {
	UserId int32 `json:"user_id" binding:"required"`
}

type GetPostsByUserIdResponse struct {
	Posts []*PostResponse `json:"posts" binding:"required"`
}

func NewGetPostsByUserIdResponse(posts []*model.Post) *GetPostsByUserIdResponse {
	postResponses := make([]*PostResponse, len(posts))
	for i, post := range posts {
		postResponses[i] = NewPostResponse(post)
	}
	return &GetPostsByUserIdResponse{Posts: postResponses}
}
