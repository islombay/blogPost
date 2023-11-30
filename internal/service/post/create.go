package post

import (
	"encoding/json"
	"errors"
	"github.com/islombay/blogPost/internal/database/postgres/post_postgres"
	"github.com/islombay/blogPost/pkg/utils/logger/sl"
	"log/slog"
)

func (s *PostService) Create(m CreateModel) (PostModel, error) {
	var r PostModel
	if !m.IsValid() {
		return r, errors.New(InvalidModel)
	}
	data, err := json.Marshal(&m)
	if err != nil {
		slog.Error("could not marshal in create service", sl.Err(err))
		return r, err
	}
	var v post_postgres.PostModel
	if err = json.Unmarshal(data, &v); err != nil {
		slog.Error("could not unmarshal in create service", sl.Err(err))
		return r, err
	}

	v, err = s.db.CreateNew(v)
	if err != nil {
		slog.Error("could not create new post in database", sl.Err(err))
		return r, err
	}

	data, err = json.Marshal(&v)
	if err != nil {
		slog.Error("could not marshal in create handler", sl.Err(err))
		return r, err
	}
	if err = json.Unmarshal(data, &r); err != nil {
		slog.Error("could not unmarshal in create handler", sl.Err(err))
		return r, err
	}
	return r, nil
}
