package request_model

import (
	"encoding/json"
	"net/http"
)

type GetUserByUsernameRequest struct {
	Username string `json:"username"`
}

func NewGetUserByUsernameRequest(r *http.Request) (*GetUserByUsernameRequest, error) {
	var request GetUserByUsernameRequest
    err := json.NewDecoder(r.Body).Decode(&request)
    return &request, err
}