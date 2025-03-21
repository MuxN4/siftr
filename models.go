package main

import (
	"database/sql"

	"github.com/MuxN4/siftr/internal/db"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	ApiKey    string       `json:"api_key"`
}

func databaseUserToUser(dbUser db.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	Url       string       `json:"url"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	UserID    uuid.UUID    `json:"user_id"`
}

func databaseFeedToFeed(dbFeed db.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []db.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollowers struct {
	ID        uuid.UUID    `json:"id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	UserID    uuid.UUID    `json:"user_id"`
	FeedID    uuid.UUID    `json:"feed_id"`
}

func databaseFeedFollowersToFeedFollowers(dbFeedFollower db.FeedFollower) FeedFollowers {
	return FeedFollowers{
		ID:        dbFeedFollower.ID,
		CreatedAt: dbFeedFollower.CreatedAt,
		UpdatedAt: dbFeedFollower.UpdatedAt,
		UserID:    dbFeedFollower.UserID,
		FeedID:    dbFeedFollower.FeedID,
	}
}

func databaseFeedsFollowersToFeedsFollowers(dbFeedsFollower []db.FeedFollower) []FeedFollowers {
	feedsFollower := []FeedFollowers{}
	for _, dbFeedsFollower := range dbFeedsFollower {
		feedsFollower = append(feedsFollower, databaseFeedFollowersToFeedFollowers(dbFeedsFollower))
	}
	return feedsFollower
}
