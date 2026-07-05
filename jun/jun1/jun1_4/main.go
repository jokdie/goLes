package main

import "fmt"

type Printer interface {
	Print()
}

type User struct {
	Name string
}

func (u User) Print() {
	fmt.Println(u.Name)
}

func GetPrinter(name string) Printer {
	if name == "" {
		return nil
	}

	return User{Name: name}
}

func main() {
	u := GetPrinter("Bob")
	if u != nil {
		u.Print()
	}

	u = GetPrinter("")
	if u != nil {
		u.Print()
	}
}
