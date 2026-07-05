package main

import (
	"fmt"

	"config/config"
)

func main() {
	if c, err := config.New("www.google.com", 8080); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("host: [%s], port: [%d]\n", c.Host(), c.Port())
	}
	if c, err := config.New("www.google.com", 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("host: [%s], port: [%d]\n", c.Host(), c.Port())
	}
	if c, err := config.New("", 8080); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("host: [%s], port: [%d]\n", c.Host(), c.Port())
	}
}
