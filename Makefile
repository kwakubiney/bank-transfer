migrate-up:
	cd internal/migrations && goose postgres "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable" up
.PHONY: migrate-up

migrate-down:
	cd internal/migrations && goose postgres "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable" down
.PHONY: migrate-down

test-migrate-up:
	cd internal/migrations && goose postgres "postgresql://postgres:postgres@localhost:6000/bank_test?sslmode=disable" up
.PHONY: test-migrate-up

test-migrate-down:
	cd internal/migrations && goose postgres "postgresql://postgres:postgres@localhost:6000/bank_test?sslmode=disable" down
.PHONY: test-migrate-down