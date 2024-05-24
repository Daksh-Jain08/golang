package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for i:=0; i<len(days); i++{
		fmt.Println(days[i]);
	}

	for i := range days{
		fmt.Println(days[i])
	}

	for _, value := range days{
		fmt.Printf("value is %v\n", value)
	}

	rougueValue := 1

	for rougueValue<10 {
		if rougueValue == 2{
			goto lco
		}
		if rougueValue==5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

	lco:
		fmt.Println("Jumping to hello")
}
