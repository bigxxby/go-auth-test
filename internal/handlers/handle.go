package handlers

import (
	"net/http"
)

func (m *MainStruct) AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		http.Error(w, "Empty id", http.StatusBadRequest)
		return
	}
	// SaveTokensToMongo

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (m *MainStruct) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.URL.Query().Get("refreshToken")
	// GetUserIDByRefreshTokenFromMongo

	// refresh in db

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
