-- File: db/migrations/20230430000001_create_users_table.up.sql
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL,
                       deleted_at TIMESTAMP,
                       username VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       password VARCHAR(255) NOT NULL
);
