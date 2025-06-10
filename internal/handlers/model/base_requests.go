package request_model

type EmptyResponse struct {}

type SuccessCreationResponse struct {
	Message string `json:"message"`
}

func NewSuccessCreationResponse(message string) *SuccessCreationResponse {
	return &SuccessCreationResponse{
		Message: message,
	}
}
