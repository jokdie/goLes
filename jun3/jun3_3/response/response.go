package response

import "encoding/json"

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SuccessResponse struct {
	Success bool         `json:"success"`
	Data    UserResponse `json:"data"`
}

func CreateSuccessResponse(user UserResponse) ([]byte, error) {
	resp := SuccessResponse{
		Success: true,
		Data:    user,
	}

	return json.Marshal(resp)
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func CreateErrorResponse(message string) ([]byte, error) {
	err := ErrorResponse{
		Success: false,
		Error:   message,
	}

	return json.Marshal(err)
}
