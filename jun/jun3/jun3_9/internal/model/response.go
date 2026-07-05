package model

type SuccessResponse struct {
	ID int `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
