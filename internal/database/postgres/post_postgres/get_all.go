package post_postgres

func (r *PostPostgres) GetAll() ([]PostModel, error) {
	q := "SELECT * FROM posts"

	res := []PostModel{}
	err := r.db.Select(&res, q)
	return res, err
}
