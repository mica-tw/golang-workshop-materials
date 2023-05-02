-- File: db/migrations/20230430000002_create_posts_table.up.sql
CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL,
                       deleted_at TIMESTAMP,
                       title VARCHAR(255) NOT NULL,
                       body TEXT NOT NULL,
                       user_id INTEGER NOT NULL REFERENCES users(id)
);