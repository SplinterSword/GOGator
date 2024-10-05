-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    feed_id UUID REFERENCES feeds(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_feed UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;