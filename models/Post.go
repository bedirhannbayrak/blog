package models

import (
	"github.com/Kamva/mgm/v2"
)

type Post struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Content          string `json:"content" bson:"content"`
}

func CreatePost(title, content string) *Post {
	return &Post{
		Title:       title,
		Content: content,
	}
}
