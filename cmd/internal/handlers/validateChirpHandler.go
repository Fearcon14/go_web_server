package handlers

import (
	"net/http"

	"encoding/json"
)

type ChirpResponse struct {
	Error string `json:"error"`
	Valid bool   `json:"valid"`
}

type ChirpRequest struct {
	Body string `json:"body"`
}

func ValidateChirpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var chirpRequest ChirpRequest
	err := decoder.Decode(&chirpRequest)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if len(chirpRequest.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	w.WriteHeader(http.StatusOK)
	respondWithJSON(w, http.StatusOK, ChirpResponse{Error: "", Valid: true})

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	data, err := json.Marshal(ChirpResponse{Error: message, Valid: false})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}
