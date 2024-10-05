-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE name = $1
LIMIT 1;

-- name: GetFeedFromURL :one
SELECT * FROM feeds
WHERE url = $1
LIMIT 1;

-- name: ListFeeds :many    
SELECT * FROM feeds;

-- name: DeleteAllFeeds :exec
DELETE FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT id, url
FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;
