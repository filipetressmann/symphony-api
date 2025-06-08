package model


type Post struct {
	PostId    int32
	UserId    int32  `json:"user_id"`
	Text      string `json:"text"`
	UrlFoto   string `json:"url_foto"`
	LikeCount int    `json:"like_count"`
}

func NewPost(
	postId int32,
	userId int32,
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
		"user_id":    post.UserId,
		"text":       post.Text,
		"url_foto":   post.UrlFoto,
		"like_count": post.LikeCount,
	}
}

func MapToPost(data map[string]any) *Post {
	return &Post{
		PostId:    data["id"].(int32),
		UserId:    data["user_id"].(int32),
		Text:      data["text"].(string),
		UrlFoto:   data["url_foto"].(string),
		LikeCount: int(data["like_count"].(int32)),
	}
}
