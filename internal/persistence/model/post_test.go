package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPost(t *testing.T) {
	p := NewPost(1, 2, "Test post", "test.jpg", 0)

	assert.NotNil(t, p)
	assert.Equal(t, int32(1), p.PostId)
	assert.Equal(t, int32(2), p.UserId)
	assert.Equal(t, "Test post", p.Text)
	assert.Equal(t, "test.jpg", p.UrlFoto)
	assert.Equal(t, 0, p.LikeCount)
}

func TestPostToMap(t *testing.T) {
	p := NewPost(1, 2, "Test post", "test.jpg", 5)

	m := p.ToMap()

	assert.Equal(t, p.UserId, m["user_id"])
	assert.Equal(t, p.Text, m["text"])
	assert.Equal(t, p.UrlFoto, m["url_foto"])
	assert.Equal(t, p.LikeCount, m["like_count"])
	// Verify id is not in the map since it's auto-generated
	_, exists := m["id"]
	assert.False(t, exists)
}

func TestMapToPost(t *testing.T) {
	data := map[string]any{
		"id":         int32(1),
		"user_id":    int32(2),
		"text":       "Test post",
		"url_foto":   "test.jpg",
		"like_count": int32(5),
	}

	p := MapToPost(data)

	assert.NotNil(t, p)
	assert.Equal(t, int32(1), p.PostId)
	assert.Equal(t, int32(2), p.UserId)
	assert.Equal(t, "Test post", p.Text)
	assert.Equal(t, "test.jpg", p.UrlFoto)
	assert.Equal(t, 5, p.LikeCount)
}
