package db

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func MigrateDB(dsn string) error {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	goose.SetBaseFS(MigrationFiles())

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "migrate"); err != nil {
		return err
	}
	return nil
}
