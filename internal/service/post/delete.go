package post

import (
	"github.com/islombay/blogPost/pkg/utils/logger/sl"
	"log/slog"
)

func (s *PostService) Delete(id int64) error {
	if err := s.db.Delete(id); err != nil {
		slog.Error("could not delete post in db", sl.Err(err))
		return err
	}
	return nil
}
