DATABASE_PASSWORD ?= postgrespassword
DATABASE_URL = postgresql://postgres:$(DATABASE_PASSWORD)@localhost:5432/ecnc_oa_db?sslmode=disable

.PHONY: run
run:
	@go run ./cmd/api

.PHONY: new_migration
new_migration:
	@migrate create -ext sql -dir db/migrations -seq $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate_up
migrate_up:
	@migrate -database $(DATABASE_URL) -path db/migrations up $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate_down
migrate_down:
	@migrate -database $(DATABASE_URL) -path db/migrations down $(filter-out $@,$(MAKECMDGOALS))

%:
	@: