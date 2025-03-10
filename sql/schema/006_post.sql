-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT ,
    published_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL UNIQUE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;