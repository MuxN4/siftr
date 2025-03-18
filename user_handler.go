package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MuxN4/siftr/internal/db"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {

	type Parameter struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := Parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		sendErrorResponse(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:   uuid.New(),
		Name: params.Name,
	})
	if err != nil {
		sendErrorResponse(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	sendJSONResponse(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) GerUserHandler(w http.ResponseWriter, r *http.Request, user db.User) {
	sendJSONResponse(w, 200, databaseUserToUser(user))
}
