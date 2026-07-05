package dto

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
