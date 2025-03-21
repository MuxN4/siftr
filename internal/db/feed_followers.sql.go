// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feed_followers.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createFeedFollowers = `-- name: CreateFeedFollowers :one
INSERT INTO feed_followers (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, DEFAULT, DEFAULT, $2, $3)
RETURNING id, created_at, updated_at, user_id, feed_id
`

type CreateFeedFollowersParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
	FeedID uuid.UUID
}

func (q *Queries) CreateFeedFollowers(ctx context.Context, arg CreateFeedFollowersParams) (FeedFollower, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollowers, arg.ID, arg.UserID, arg.FeedID)
	var i FeedFollower
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const deleteFeedFollowers = `-- name: DeleteFeedFollowers :exec
DELETE FROM feed_followers WHERE id=$1 AND user_id=$2
`

type DeleteFeedFollowersParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteFeedFollowers(ctx context.Context, arg DeleteFeedFollowersParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollowers, arg.ID, arg.UserID)
	return err
}

const getFeedFollowers = `-- name: GetFeedFollowers :many
SELECT id, created_at, updated_at, user_id, feed_id FROM feed_followers WHERE user_id=$1
`

func (q *Queries) GetFeedFollowers(ctx context.Context, userID uuid.UUID) ([]FeedFollower, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowers, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollower
	for rows.Next() {
		var i FeedFollower
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
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
