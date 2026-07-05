package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "my string"
	}()

	str := <-messages

	fmt.Println(str)
}
