package main

import "fmt"

type Writer interface {
	Write(text string)
}

type File struct {
	Content string
}

func (f *File) Write(text string) {
	f.Content += text
}

func SaveHello(w Writer) {
	w.Write("hello")
}

func main() {
	f := &File{}

	SaveHello(f)

	fmt.Println(f.Content)
}
