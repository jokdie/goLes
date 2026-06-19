package main

import (
	"fmt"
	"jun2_5/storage"
)

func printUserByID(s storage.Storage, id int) {
	if u, err := s.GetByID(id); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ID: %d, Name: %s\n", u.ID, u.Name)
	}
}

func main() {
	s := storage.New()

	printUserByID(s, 1)
	printUserByID(s, 3)
}
