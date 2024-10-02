-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE name = $1
LIMIT 1;

-- name: ListFeeds :many    
SELECT * FROM feeds;

-- name: DeleteFeed :exec
DELETE FROM feeds
WHERE name = $1;