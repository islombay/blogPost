package postgres

func (db *PostgresDB) Close() error {
	return db.db.Close()
}
