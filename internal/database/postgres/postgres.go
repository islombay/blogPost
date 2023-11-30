package postgres

import (
	"fmt"
	"github.com/islombay/blogPost/internal/database/postgres/post_postgres"
	"github.com/islombay/blogPost/pkg/utils/logger/sl"
	"log/slog"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
)

type PostgresDB struct {
	db   *sqlx.DB
	Post post_postgres.PostInterface
}

type NewPostgresDBBody struct {
	Host     string
	Port     string
	Username string
	PWD      string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(a NewPostgresDBBody) *PostgresDB {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", a.Username, a.PWD, a.Host, a.Port, a.DBName, a.SSLMode),
	)
	if err != nil {
		slog.Error("could not create database connection", sl.Err(err))
		os.Exit(1)
	}
	if err = db.Ping(); err != nil {
		slog.Error("could not ping database", sl.Err(err))
	}

	if err = makeMigrations(a.Host, a.Port, a.Username, a.PWD, a.DBName, a.SSLMode); err != nil && err != migrate.ErrNoChange {
		slog.Error("could not migrate up", sl.Err(err))
		os.Exit(1)
	}
	slog.Info("migrations are applied")

	return &PostgresDB{
		db:   db,
		Post: post_postgres.NewPostPostgres(db),
	}
}

func makeMigrations(host, port, user, pwd, dbname, ssl string) error {
	m, err := migrate.New("file://internal/database/postgres/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, pwd, host, port, dbname, ssl))

	if err != nil {
		slog.Error("could not open migrations", sl.Err(err))
		os.Exit(1)
	}

	return m.Up()
}
