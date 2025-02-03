package db

import "embed"

//go:embed migrate/*.sql
var embededMigrations embed.FS

func MigrationFiles() embed.FS {
	return embededMigrations
}
