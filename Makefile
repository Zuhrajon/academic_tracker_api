include utils.mk

.PHONY: up
up:
	docker compose up --force-recreate --build -d
	$(MAKE) migration-up

.PHONY: generate-migration
# usage: make generate-migration name=create_students_table
generate-migration:
	goose -dir ./migrations create $(name) sql

.PHONY: mocks
mocks:
	@echo "mocks generation is not implemented yet"

.PHONY: migration-up
migration-up:
	goose -dir ./migrations postgres "$(academic_tracker_postgres_dsn)" up

.PHONY: migration-down
migration-down:
	goose -dir ./migrations postgres "$(academic_tracker_postgres_dsn)" down-to 0

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -count=1 ./...

.PHONY: test-int
test-int:
	go test -count=1 ./tests/integration/... --tags=integration

.PHONY: check
check: lint test test-int

.PHONY: down
down:
	docker compose down -v