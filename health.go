package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, 200, struct{}{})
}
