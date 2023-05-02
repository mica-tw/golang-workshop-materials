-- File: db/migrations/20230430000005_create_post_tags_table.up.sql
CREATE TABLE post_tags (
                           post_id INTEGER NOT NULL REFERENCES posts(id),
                           tag_id INTEGER NOT NULL REFERENCES tags(id),
                           PRIMARY KEY (post_id, tag_id)
);