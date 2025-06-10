package base_handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func MapRequest[T any](r *http.Request) (*T, error) {
	request := new(T)
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		log.Printf("Request body: %v", r.Body)
		return nil, err
	}

	return request, nil
}

func MustEncodeAnswer(data any, w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Printf("Error encoding answer: %s", err)
		http.Error(w, "Error while encoding answer", http.StatusInternalServerError)
		return
	}
}
