-- name: CreateFeedFollow :many
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id,created_at,updated_at,feed_id, user_id) VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
JOIN users ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows
JOIN feeds ON feed_follows.feed_id = feeds.id
JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING feeds
WHERE feed_follows.user_id = $1
AND feeds.url = $2
AND feed_follows.feed_id = $3;


-- name: DeleteAllFeedFollows :exec
DELETE FROM feed_follows;