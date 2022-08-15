# bank-transfer

A REST API for a simple core banking system.

# Setting up

- Create an `.env` file in root directory by cloning `.env_test` (change `POSTGRES_DB` to `bank` & `POSTGRES_HOST` to `db` in `.env` file)

- Run `docker-compose up` to spin up web server, development database and test database.

- Run development database migrations with `goose postgres "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable" up` in the `internal/migrations` directory
  
- Run test database migrations with `goose postgres "postgresql://postgres:postgres@localhost:6000/bank_test?sslmode=disable" up` in the `internal/migrations` directory
  
- Check out APIDocs next ;)

# TODO

- Make database transactions across different repositories. Example: Writes to Accounts and Transaction repositories must be atomic to prevent inconsistencies.

- Include retries with idempotency.
