package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to function in golang!")

	proResult, message := proAdder(3,5,4,6);
	result := adder(3,5);
	
	fmt.Println("Result is: ", result)
	fmt.Println("proResult is: ", proResult)
	fmt.Println("Message is: ", message)
}

func adder(a int, b int) int {
	return a+b;
}

func greeter() {
	fmt.Println("Namaste from golang")
}

func proAdder(values ...int) (int, string) {
	total := 0;

	for _, val := range values {
		total += val
	}

	return total, "Hello i am proAdder function"
}