package main

import (
	"fmt"
	"jun3_3/response"
)

func main() {
	u := response.UserResponse{
		ID:    1,
		Name:  "Bill",
		Email: "progib322@hq.com",
	}

	if res, err := response.CreateSuccessResponse(u); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(string(res))
	}

	if res, err := response.CreateErrorResponse("user not found"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(string(res))
	}
}
