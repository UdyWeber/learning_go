package main

import "fmt"

// Swapping numerical variables in Go
func GoSwap() {
	x := 10
	y := 70

	fmt.Printf("X was %d, y was %d\n", x, y)

	x = x + y
	y = x - y
	x = x - y

	fmt.Printf("Now X is %d and y is %d\n", x, y)

}

// Printout tests that I've made with string prints
func StringTests() {
	bytes, bytes_error := fmt.Println("Hello World!",)
	
	if bytes_error != nil {
		fmt.Println("Error: Deu ruim manéé")
	}
	
	fmt.Println("Number of bytes:", bytes)
}

// Printout the primitive types used in a formatted String
func PrimitiveTypes() {
	
	a_boolean := false
	a_number := 102
	a_full_string := "Today is a happy day"

	fmt.Printf("(day %d) %s: %t\n", a_number, a_full_string, a_boolean)
	fmt.Printf("Types for each variables: %T, %T, %T\n", a_number, a_full_string, a_boolean)
}
