package dtos

import (
	"first-go-api/models"
	"time"
)

type Post struct {
	Title  string `json:"title" binding:"required" example:"First Post"`
	Body   string `json:"body" binding:"required" example:"This is the body of the first post."`
	Likes  int    `json:"likes" example:"100"`
	Draft  bool   `json:"draft" example:"false"`
	Author string `json:"author" binding:"required" example:"John Doe"`
}

type PostResponse struct {
	ID        uint      `json:"id" example:"1"`
	Title     string    `json:"title" example:"First Post"`
	Body      string    `json:"body" example:"This is the body of the first post."`
	Likes     int       `json:"likes" example:"100"`
	Draft     bool      `json:"draft" example:"false"`
	Author    string    `json:"author" example:"John Doe"`
	CreatedAt time.Time `json:"created_at" example:"2024-10-06T10:32:36Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-10-06T10:32:36Z"`
}

// NewPostResponse é o "construtor" que encapsula um models.Post em um PostResponse
func NewPostResponse(post models.Post) PostResponse {
	return PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		Likes:     post.Likes,
		Draft:     post.Draft,
		Author:    post.Author,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

// NewPostResponseList converte uma lista de models.Post em uma lista de PostResponse
func NewPostResponseList(posts []models.Post) []PostResponse {
	postResponses := make([]PostResponse, len(posts))
	for i, post := range posts {
		postResponses[i] = NewPostResponse(post)
	}
	return postResponses
}
