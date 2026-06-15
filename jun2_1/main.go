package main

import (
	"fmt"

	"user/user"
)

func main() {
	u := user.New(1, "Arkadiy")

	fmt.Printf("ID: %d, Name: %s", u.ID, u.Name)
}
