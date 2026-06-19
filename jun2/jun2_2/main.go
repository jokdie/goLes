package main

import (
	"fmt"
	"user/user"
)

func main() {
	u := user.New(1, "Boris")

	fmt.Printf("id: %d, name: %s", u.ID(), u.Name())
}
