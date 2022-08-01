# bank-transfer

A REST API for a simple core banking system.

# Setting up

- Create an `.env` file in root directory by cloning `.env_test`
- 
- Run `docker-compose` up to spin up web server and database

- Run database migrations with `goose postgres "user=postgres password=postgres dbname=bank sslmode=disable" up` in the `internal/migrations` directory

- Check out APIDocs next ;)