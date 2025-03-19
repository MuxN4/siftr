package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MuxN4/siftr/internal/db"
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
