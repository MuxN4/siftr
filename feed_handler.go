package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MuxN4/siftr/internal/db"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user db.User) {

	type Parameter struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := Parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		sendErrorResponse(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID:     uuid.New(),
		Name:   params.Name,
		Url:    params.URL,
		UserID: user.ID,
	})
	if err != nil {
		sendErrorResponse(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	sendJSONResponse(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		sendErrorResponse(w, 400, fmt.Sprintf("Error getting feed: %v", err))
		return
	}

	sendJSONResponse(w, 201, databaseFeedsToFeeds(feeds))
}
