-- File: db/migrations/20230430000003_create_comments_table.up.sql
CREATE TABLE comments (
                          id SERIAL PRIMARY KEY,
                          created_at TIMESTAMP NOT NULL,
                          updated_at TIMESTAMP NOT NULL,
                          deleted_at TIMESTAMP,
                          body TEXT NOT NULL,
                          user_id INTEGER NOT NULL REFERENCES users(id),
                          post_id INTEGER NOT NULL REFERENCES posts(id)
);