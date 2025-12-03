package handlers

import (
	"net/http"

	"encoding/json"
)

func ValidateChirpHandler(w http.ResponseWriter, r *http.Request) {
	type ChirpRequest struct {
		Body string `json:"body"`
	}

	type ChirpResponse struct {
		Error string `json:"error"`
		Valid bool   `json:"valid"`
	}

	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var chirpRequest ChirpRequest
	err := decoder.Decode(&chirpRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, err := json.Marshal(ChirpResponse{Error: "Invalid request body", Valid: false})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	if len(chirpRequest.Body) > 140 {
		w.WriteHeader(http.StatusBadRequest)
		data, err := json.Marshal(ChirpResponse{Error: "Chirp is too long", Valid: false})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(ChirpResponse{Error: "", Valid: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)

}
