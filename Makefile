POSTGRES_URI=postgresql://yaroslav:AnySecretPassword!!@localhost:5432/yaaws?sslmode=disable

env-up:
	docker-compose up -d

env-down:
	docker-compose down --remove-orphans -v

go-test:
	docker exec research-sqlc-go-app go test ./... -v -count=1

docker-go-version:
	docker exec research-sqlc-go-app go version

docker-pg-version:
	docker exec research-sqlc-postgres psql -U yaroslav -d yaaws -c "SELECT VERSION();"

test:
	make env-up
	make docker-go-version
	make docker-pg-version
	make migrate-up
	make go-test
	make env-down

# go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# sqlc generate
#
# alternative
# docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate
generate-sqlc:
	sqlc generate

# Creates new migration file with the current timestamp
# Example: make create-new-migration-file NAME=<name>
create-new-migration-file:
	$(eval NAME ?= noname)
	mkdir -p ./internal/storage/postgres/migrations/
	goose -dir ./internal/storage/postgres/migrations/ create $(NAME) sql

migrate-up:
	goose -dir ./internal/storage/postgres/migrations/ -table schema_migrations postgres $(POSTGRES_URI) up
migrate-redo:
	goose -dir ./internal/storage/postgres/migrations/ -table schema_migrations postgres $(POSTGRES_URI) redo
migrate-down:
	goose -dir ./internal/storage/postgres/migrations/ -table schema_migrations postgres $(POSTGRES_URI) down
migrate-reset:
	goose -dir ./internal/storage/postgres/migrations/ -table schema_migrations postgres $(POSTGRES_URI) reset
migrate-status:
	goose -dir ./internal/storage/postgres/migrations/ -table schema_migrations postgres $(POSTGRES_URI) status
