package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "Hello i am daksh jain."

	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is", length)
	defer file.Close()
	readFile("./file.txt")
}

func readFile(filename string) {
	databytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("The text inside the file:\n", string(databytes))
}
