package models

type Post struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Content     string `json:"content"`
	IsPublished bool   `json:"isPublished"`
}
