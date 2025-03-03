// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeeds = `-- name: CreateFeeds :one
INSERT INTO feeds (id, create_at, update_at, name, url , user_id )
VALUES ($1, $2, $3, $4 , $5, $6)
RETURNING id, create_at, update_at, name, url, user_id, last_fetch_at
`

type CreateFeedsParams struct {
	ID       uuid.UUID
	CreateAt time.Time
	UpdateAt time.Time
	Name     string
	Url      string
	UserID   uuid.UUID
}

func (q *Queries) CreateFeeds(ctx context.Context, arg CreateFeedsParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeeds,
		arg.ID,
		arg.CreateAt,
		arg.UpdateAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreateAt,
		&i.UpdateAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchAt,
	)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT id, create_at, update_at, name, url, user_id, last_fetch_at FROM feeds
`

func (q *Queries) GetFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreateAt,
			&i.UpdateAt,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.LastFetchAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNextFeedToFetch = `-- name: GetNextFeedToFetch :many
SELECT id, create_at, update_at, name, url, user_id, last_fetch_at
FROM feeds
ORDER BY last_fetch_at ASC NULLS FIRST
LIMIT $1
`

func (q *Queries) GetNextFeedToFetch(ctx context.Context, limit int32) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getNextFeedToFetch, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreateAt,
			&i.UpdateAt,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.LastFetchAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const markFeedAsFetched = `-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetch_at = NOW() , update_at = NOW()
WHERE id = $1
RETURNING id, create_at, update_at, name, url, user_id, last_fetch_at
`

func (q *Queries) MarkFeedAsFetched(ctx context.Context, id uuid.UUID) (Feed, error) {
	row := q.db.QueryRowContext(ctx, markFeedAsFetched, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreateAt,
		&i.UpdateAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchAt,
	)
	return i, err
}
