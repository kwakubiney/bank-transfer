-- +goose Up
CREATE TABLE IF NOT EXISTS transactions
(
	id 				UUID 		PRIMARY KEY DEFAULT gen_random_uuid(),	
	credit          VARCHAR(36)      	NOT NULL,
	debit           VARCHAR(36),
    amount          BIGINT      NOT NULL,
	created_at 		TIMESTAMP 	NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS transactions;
