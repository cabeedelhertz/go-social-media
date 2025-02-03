
server:
	go run ./cmd/social/... server



db-create:
	psql -d postgres -c "CREATE DATABASE $(SERVICE_NAME)_local"

db-drop:
	psql -d postgres -c "DROP DATABASE $(SERVICE_NAME)_local"
	
migrate-up:
	goose -dir "./db/migrate" postgres 'user=${PGUSER} dbname=${PGDATABASE} sslmode=disable password=${PGPASSWORD}' up

migrate-down:
	goose -dir "./db/migrate" postgres 'user=${PGUSER} dbname=${PGDATABASE} sslmode=disable password=${PGPASSWORD}' down

db-init: db-drop db-create migrate-up