package handlers

import (
	"fmt"
	"net/http"
)

// RootHandler is the root handler for the API.
//	@Summary		Root Handler
//	@Description	Returns a simple greeting message.
//	@Tags			Root
//	@Produce		plain
//	@Success		200	{string}	string	"Olá, mundo!"
func RootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, "Hello, World!"); err != nil {
			http.Error(w, "Failed in writing answer.", http.StatusInternalServerError)
		}
	}
}
