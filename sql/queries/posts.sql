/*
    id UUID PRIMARY KEY,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT ,
    published_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL UNIQUE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
*/ 
-- name: CreatePost :one
INSERT INTO posts (id, create_at,update_at,title,description,published_at,url,feed_id)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;


/*
    will fetch feed only user follow so we join table what post it's belong feed user follow by
    join feed_follows to post table
*/
-- name: GetPostForUser :many
SELECT posts.* from posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;