package post_postgres

import (
	"time"
)

type PostModel struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Username  *string   `json:"username"`
}
