package handlers

import (
	"net/http"

	"encoding/json"
	"strings"
)

type ChirpResponse struct {
	Error        string `json:"error"`
	Cleaned_Body string `json:"cleaned_body"`
}

type ChirpRequest struct {
	Body string `json:"body"`
}

var profanityList = []string{"kerfuffle", "sharbert", "fornax"}

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

	cleanedBody := removeProfanity(chirpRequest.Body)
	respondWithJSON(w, http.StatusOK, ChirpResponse{Error: "", Cleaned_Body: cleanedBody})

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	data, err := json.Marshal(ChirpResponse{Error: message, Cleaned_Body: ""})
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

func removeProfanity(body string) string {
	splitBody := strings.Split(body, " ")
	for i, word := range splitBody {
		for _, profanity := range profanityList {
			if strings.ToLower(word) == strings.ToLower(profanity) {
				splitBody[i] = "****"
			}
		}
	}
	return strings.Join(splitBody, " ")
}
