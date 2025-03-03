-- name: CreateFeeds :one
INSERT INTO feeds (id, create_at, update_at, name, url , user_id )
VALUES ($1, $2, $3, $4 , $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds ;

-- name: GetNextFeedToFetch :many
SELECT *
FROM feeds
ORDER BY last_fetch_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one 
UPDATE feeds
SET last_fetch_at = NOW() , update_at = NOW()
WHERE id = $1
RETURNING *;