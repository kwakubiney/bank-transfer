-- +goose Up
CREATE TABLE IF NOT EXISTS accounts
(
	id 				UUID 		PRIMARY KEY DEFAULT gen_random_uuid(),	
	name         	VARCHAR(50) NOT NULL,
	balance			INTEGER 	NOT NULL,
	created_at 		TIMESTAMP 	NOT NULL,
	last_modified   TIMESTAMP   NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS accounts;
