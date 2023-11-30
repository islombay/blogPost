package post

import (
	"github.com/islombay/blogPost/internal/database/postgres/post_postgres"
)

const (
	InvalidModel         = "create_model_invalid"
	ErrorMarshalObject   = "could_not_marshal"
	ErrorUnmarshalObject = "could_not_unmarshal"
)

//go:generate mockgen -source=post.go -destination=mocks/mock.go

type PostServiceInterface interface {
	Create(m CreateModel) (PostModel, error)
	GetAll() ([]PostModel, error)
	Delete(id int64) error
}

type PostService struct {
	db post_postgres.PostInterface
}

func NewPostService(postInterface post_postgres.PostInterface) PostServiceInterface {
	return &PostService{db: postInterface}
}
