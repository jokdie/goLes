package main

import (
	"encoding/json"
	"fmt"
	"jun3_2/user"
)

func main() {
	u := user.User{
		ID:    1,
		Name:  "Grisha",
		Email: "nagibator67rus@mail.ru",
	}

	jsonData, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(string(jsonData))

	var restored user.User

	err = json.Unmarshal(jsonData, &restored)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(restored)
}
