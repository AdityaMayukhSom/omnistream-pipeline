package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Content     string `json:"content"`
	IsPublished bool   `json:"isPublished"`
}
