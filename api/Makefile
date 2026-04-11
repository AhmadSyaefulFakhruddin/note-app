DB_URL=postgres://postgres:123qweas@localhost:5432/note_db?sslmode=disable

migrate-up:
	migrate -path db/migrations -database $(DB_URL) up

migrate-down:
	migrate -path db/migrations -database $(DB_URL) down

migrate-status:
	migrate -path db/migrations -database $(DB_URL) version

migrate-force:
	migrate -path db/migrations -database $(DB_URL) force 1