package model

type Tag struct {
	TagID   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

type Article struct {
	ArticleID     int    `json:"article_id"`
	ArticleTitle  string `json:"article_title"`
	ArticleUserID int    `json:"article_user_id"`
	ArticleBody   string `json:"article_body"`
	ImageBase64   string `json:"image_base_64"`
}

type ArticleTag struct {
	ArticleTagID int `json:"posttag_id`
	ArticleID    int `json:"posttag_post_id"`
	TagID        int `json:"posttag_tag_id"`
}
