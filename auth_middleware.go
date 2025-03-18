package main

import (
	"fmt"
	"net/http"

	"github.com/MuxN4/siftr/internal/auth"
	"github.com/MuxN4/siftr/internal/db"
)

type authedHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			sendErrorResponse(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		// Fetch user from the database
		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			sendErrorResponse(w, http.StatusNotFound, "User not found")
			return
		}

		handler(w, r, user)
	}
}
