package service

import (
	"github.com/islombay/blogPost/internal/database/postgres"
	"github.com/islombay/blogPost/internal/service/post"
)

type BlogPostService struct {
	db   *postgres.PostgresDB
	Post post.PostServiceInterface
}

func NewBlogPostService(db *postgres.PostgresDB) *BlogPostService {
	return &BlogPostService{
		db:   db,
		Post: post.NewPostService(db.Post),
	}
}
