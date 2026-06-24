package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	User User `json:"user"`
}
