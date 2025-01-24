## migrate: Run the migration to migrate the tables in the database
migrate:
	go run src/infrastructure/database/postgres/migration/main.go

## run: Run the application
run:
	go run src/infrastructure/gofiber/main.go

## build: Build the application
build:
	go build src/infrastructure/gofiber/main.go

.PHONY: migrate run build
