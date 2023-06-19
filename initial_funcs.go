package main

import "fmt"

// Swapping numerical variables in Go
func GoSwap() {
	x, y := 10, 70

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
// Variables declared outside of codeblocks must be var {name} = {value}
var some_var int = 10

func PritnPackageVariable() {
	fmt.Println("Package variable:", some_var)
}

// Function that shows empty return of primitive types
func ZeroValuesForTypes() {
	var empty_integer int
	var empty_float float64
	var empty_boolean bool
	var empty_string string

	fmt.Printf(
		"Empty Integer: %v.\nEmpty Float: %v.\nEmpty Boolean: %v.\nEmpty String: %v.\nEmpty Function: %v.\n",
		empty_integer, 
		empty_float, 
		empty_boolean, 
		empty_string,
		nil,
	)
}

func TypesForStrings() {
	// Interpreted String Literal
	interpreted_string := "Abluble\nWadawada?\tSopinha da \"Vovo\""
	println("Interpreted String: ", interpreted_string)

	raw_string := `ASCII Chars
		\n Blabla

			Claudeti

		Azulejo
			Golang


					maYcaKO
	`
	println("Raw string:", raw_string)
}

func InitialFunctionsRunner() {
	StringTests()
	GoSwap()
	PrimitiveTypes()
	PritnPackageVariable()
	ZeroValuesForTypes()
	TypesForStrings()
}