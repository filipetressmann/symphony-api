package handlers

import (
	"net/http"
	request_model "symphony-api/internal/handlers/model"
)

func createHandler[In any, Out any](
	handler func(In) (Out, error),
	) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := request_model.MapRequest[In](r)

		if err != nil {
        	http.Error(w, "Invalid Input", http.StatusBadRequest)
        	return
    	}

		response, err := handler(*request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		request_model.MustEncodeAnswer(response, w)
	}
}