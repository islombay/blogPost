package post_postgres

func (r *PostPostgres) Delete(id int64) error {
	q := "delete from posts where id = $1"

	_, err := r.db.Exec(q, id)
	return err
}
