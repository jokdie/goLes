package response

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
	Data    User `json:"data"`
}

func CreateSuccessResponse() SuccessResponse {
	return SuccessResponse{
		Success: true,
		Data: User{
			ID:    1,
			Name:  "Alice",
			Email: "alice@example.com",
		},
	}
}
