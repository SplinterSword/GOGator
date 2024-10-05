-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    published_at TIMESTAMPTZ NOT NULL,
    feed_id UUID REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts