package post

import (
	"encoding/json"
	"github.com/islombay/blogPost/pkg/utils/logger/sl"
	"log/slog"
)

func (s *PostService) GetAll() ([]PostModel, error) {
	v := []PostModel{}
	res, err := s.db.GetAll()
	if err != nil {
		slog.Error("error occured while getting all posts", sl.Err(err))
		return v, err
	}
	data, err := json.Marshal(res)
	if err != nil {
		slog.Error("could not marshal in get all posts", sl.Err(err))
		return v, err
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		slog.Error("could not unmarshal in get all posts", sl.Err(err))
		return v, err
	}

	return v, nil
}
