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

func (repository *PostRepository) get(constraint map[string]any) []*model.Post {
	data := repository.connection.Get(constraint, POST_TABLE)

	posts := make([]*model.Post, 0)

	for _, post := range data {
		posts = append(posts, model.MapToPost(post))
	}

	return posts
}

func (repository *PostRepository) GetById(postId int64) *model.Post {
	constraint := map[string]any{
		"post_id": postId,
	}

	posts := repository.get(constraint)
	if len(posts) == 0 {
		return nil
	}
	return posts[0]
}

func (repository *PostRepository) GetByUserId(userId int64) []*model.Post {
	constraint := map[string]any{
		"user_id": userId,
	}

	return repository.get(constraint)
}
