package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (dog Dog) Speak() string {
	return "woof"
}

type Cat struct{}

func (cat Cat) Speak() string {
	return "meow"
}

func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{}
	MakeSpeak(dog)

	cat := Cat{}
	MakeSpeak(cat)
}
