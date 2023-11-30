package post_postgres

import (
	"database/sql"
	"strings"
)

func (db *PostPostgres) CreateNew(m PostModel) (PostModel, error) {
	q := `INSERT INTO posts (title, content, created_at, username) values($1, $2, $3, $4)`

	username := sql.NullString{
		String: *m.Username,
		Valid:  true,
	}
	if strings.TrimSpace(username.String) == "" {
		username.Valid = false
	}
	_, err := db.db.Exec(q, m.Title, m.Content, m.CreatedAt, username)
	if err != nil {
		return m, err
	}
	err = db.db.Get(&m, "select * from posts where title = $1 and content = $2 and created_at = $3", m.Title, m.Content, m.CreatedAt)
	return m, err
}
