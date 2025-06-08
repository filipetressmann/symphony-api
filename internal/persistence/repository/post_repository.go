package repository

import (
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/model"
)

var POST_TABLE = "post"

type PostRepository struct {
	connection postgres.PostgreConnection
}

func NewPostRepository(connection postgres.PostgreConnection) *PostRepository {
	return &PostRepository{
		connection: connection,
	}
}

func (repository *PostRepository) Put(post *model.Post) (*model.Post, error) {
	id, err := repository.connection.Put(post.ToMap(), POST_TABLE)
	return model.NewPost(
		id,
		post.UserId,
		post.Text,
		post.UrlFoto,
		post.LikeCount,
	), err
}

func (repository *PostRepository) get(constraint map[string]any) ([]*model.Post, error) {
	data, err := repository.connection.Get(constraint, POST_TABLE)

	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)

	for _, post := range data {
		posts = append(posts, model.MapToPost(post))
	}

	return posts, nil
}

func (repository *PostRepository) GetById(postId int32) (*model.Post, error) {
	constraint := map[string]any{
		"id": postId,
	}

	posts, err := repository.get(constraint)
	if len(posts) == 0 {
		return nil, err
	}
	return posts[0], err
}

func (repository *PostRepository) GetByUserId(userId int32) ([]*model.Post, error) {
	constraint := map[string]any{
		"user_id": userId,
	}

	return repository.get(constraint)
}
