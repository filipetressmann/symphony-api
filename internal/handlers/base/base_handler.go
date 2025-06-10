package base_handlers

import (
	"net/http"
)

func CreateHandler[In any, Out any](
	handler func(In) (Out, error),
	) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := MapRequest[In](r)

		if err != nil {
        	http.Error(w, "Invalid Input", http.StatusBadRequest)
        	return
    	}

		response, err := handler(*request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		MustEncodeAnswer(response, w)
	}
}