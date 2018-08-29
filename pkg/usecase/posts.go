package usecase

import (
	"fmt"

	"punyasaya/pkg/domain"
	"punyasaya/pkg/lib/config"
)

type User struct {
	UserID   int64  `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"user_password"`
}

type Tags struct {
	TagID   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

type Posts struct {
	PostID      int    `json:"post_id"`
	PostTitle   string `json:"post_title"`
	PostUserID  int    `json:"post_user_id"`
	PostBody    string `json:"post_body"`
	ImageBase64 string `json:"image_base_64"`
}

type PostTag struct {
	PostTagID int `json:"posttag_id`
	PostID int `json:"posttag_post_id"`
	TagID int `json:"posttag_tag_id"`
}

