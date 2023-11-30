package post

import (
	"strings"
	"time"
)

type CreateModel struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

func (model *CreateModel) IsValid() bool {
	if strings.Trim(model.Title, " ") == "" {
		return false
	}
	if strings.Trim(model.Content, " ") == "" {
		return false
	}

	//now := time.Now()
	//diff := now.Sub(model.CreatedAt).Hours()
	//if diff/24 > 1 {
	//	return false
	//}
	return true
}

type PostModel struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}
