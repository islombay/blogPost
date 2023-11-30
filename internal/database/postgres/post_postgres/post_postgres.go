package post_postgres

import (
	"github.com/jmoiron/sqlx"
)

type PostInterface interface {
	CreateNew(PostModel) (PostModel, error)
	GetAll() ([]PostModel, error)
	Delete(id int64) error
}

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) PostInterface {
	return &PostPostgres{db: db}
}
