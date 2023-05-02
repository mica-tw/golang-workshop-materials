-- File: db/migrations/20230430000004_create_tags_table.up.sql
CREATE TABLE tags (
                      id SERIAL PRIMARY KEY,
                      created_at TIMESTAMP NOT NULL,
                      updated_at TIMESTAMP NOT NULL,
                      deleted_at TIMESTAMP,
                      name VARCHAR(255) NOT NULL
);