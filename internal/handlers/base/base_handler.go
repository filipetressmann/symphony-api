package base_handlers

import (
	"net/http"
)

func CreatePostMethodHandler[In any, Out any](handler func(In) (Out, error),) http.HandlerFunc {
	return createHandler(handler, MapRequest)
}

func CreateGetMethodHandler[In any, Out any](handler func(In) (Out, error)) http.HandlerFunc {
	return createHandler(handler, MapUrlValues)
}

func createHandler[In any, Out any](
	handler func(In) (Out, error),
	mapper func(*http.Request) (*In, error),
	) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := mapper(r)

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