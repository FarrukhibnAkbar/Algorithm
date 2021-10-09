package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter a number: ")
	fmt.Scan(&a)

	if a % 5 == 0 {
		fmt.Print("Fizz\n")
	}

	if a % 3 == 0 {
		fmt.Println("Buzz")
	}

}