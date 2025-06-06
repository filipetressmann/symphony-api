package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func mustEncodeAnswer(data map[string]any, w http.ResponseWriter, errorMessage string) {
	err := json.NewEncoder(w).Encode(data)
	
	if err != nil {
		log.Printf("Error enconding answer: %s", err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
        return
	}
}