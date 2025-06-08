package repository

import (
	"testing"

	"symphony-api/internal/persistence/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostRepository_Put(t *testing.T) {
	mockConn := &MockPostgreConnection{}
	repo := NewPostRepository(mockConn)

	post := &model.Post{
		UserId:    1,
		Text:      "Test post",
		UrlFoto:   "test.jpg",
		LikeCount: 0,
	}

	mockConn.On("PutReturningId", mock.Anything, POST_TABLE).Return(int32(1), nil)

	result, _ := repo.Put(post)

	assert.Equal(t, int32(1), result.PostId)
	mockConn.AssertExpectations(t)
}

func getPostTestData() (*model.Post, map[string]any) {
	post := &model.Post{
		PostId:    1,
		UserId:    1,
		Text:      "Test post",
		UrlFoto:   "test.jpg",
		LikeCount: 0,
	}
	postMap := map[string]any{
		"id":         int32(1),
		"user_id":    int32(1),
		"text":       "Test post",
		"url_foto":   "test.jpg",
		"like_count": int32(0),
	}

	return post, postMap
}

func TestPostRepository_GetById(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := NewPostRepository(mockConn)

	post, postMap := getPostTestData()

	mockConn.On("Get", mock.Anything, POST_TABLE).Return([]map[string]any{postMap}, nil)

	result, _ := repo.GetById(1)

	assert.Equal(t, post, result)
	mockConn.AssertExpectations(t)
}

func TestPostRepository_GetByUserId(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := NewPostRepository(mockConn)

	post, postMap := getPostTestData()

	mockConn.On("Get", mock.Anything, POST_TABLE).Return([]map[string]any{postMap}, nil)

	result, _ := repo.GetByUserId(1)

	assert.Equal(t, []*model.Post{post}, result)
	mockConn.AssertExpectations(t)
}
