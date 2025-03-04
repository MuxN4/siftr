package main

import "net/http"

func errorHandler(w http.ResponseWriter, r *http.Request) {
	sendErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred.")
}
