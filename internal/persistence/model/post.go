package model

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	PostId    int64
	UserId    int64  `json:"user_id"`
	Text      string `json:"text"`
	UrlFoto   string `json:"url_foto"`
	LikeCount int    `json:"like_count"`
}

func NewPost(
	postId int64,
	userId int64,
	text string,
	urlFoto string,
	likeCount int,
) *Post {
	return &Post{
		PostId:    postId,
		UserId:    userId,
		Text:      text,
		UrlFoto:   urlFoto,
		LikeCount: likeCount,
	}
}

func (post *Post) ToMap() map[string]any {
	return map[string]any{
		"post_id":    post.PostId,
		"user_id":    post.UserId,
		"text":       post.Text,
		"url_foto":   post.UrlFoto,
		"like_count": post.LikeCount,
	}
}

func PostFromRequest(request *http.Request) (*Post, error) {
	var post Post
	err := json.NewDecoder(request.Body).Decode(&post)
	return &post, err
}

func MapToPost(data map[string]any) *Post {
	return &Post{
		PostId:    data["post_id"].(int64),
		UserId:    data["user_id"].(int64),
		Text:      data["text"].(string),
		UrlFoto:   data["url_foto"].(string),
		LikeCount: int(data["like_count"].(int64)),
	}
}
