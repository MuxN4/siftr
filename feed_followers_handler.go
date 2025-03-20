package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MuxN4/siftr/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createFeedFollowersHandler(w http.ResponseWriter, r *http.Request, user db.User) {

	type Parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := Parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		sendErrorResponse(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	FeedFollowers, err := apiCfg.DB.CreateFeedFollowers(r.Context(), db.CreateFeedFollowersParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		sendErrorResponse(w, 400, fmt.Sprintf("Error creating feed followers: %v", err))
		return
	}

	sendJSONResponse(w, 201, databaseFeedFollowersToFeedFollowers(FeedFollowers))
}

func (apiCfg *apiConfig) getFeedFollowersHandler(w http.ResponseWriter, r *http.Request, user db.User) {
	FeedFollowers, err := apiCfg.DB.GetFeedFollowers(r.Context(), user.ID)
	if err != nil {
		sendErrorResponse(w, 400, fmt.Sprintf("Error getting feed followers: %v", err))
		return
	}

	sendJSONResponse(w, 201, databaseFeedsFollowersToFeedsFollowers(FeedFollowers))
}

func (apiCfg *apiConfig) deleteFeedFollowersHandler(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		sendErrorResponse(w, 400, fmt.Sprintf("Error parsing feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollowers(r.Context(), db.DeleteFeedFollowersParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		sendErrorResponse(w, 400, fmt.Sprintf("Error deleting feed followers: %v", err))
		return
	}

	sendJSONResponse(w, 200, struct{}{})

}
