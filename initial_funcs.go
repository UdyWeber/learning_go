package main

import (
	"fmt"
	"runtime"
)

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

func GetStringBytes(some_string string) {
	x := []byte(some_string)
	fmt.Printf("Bytes for %v are: %v\n", some_string, x)

	for _, letter := range x {
		// Printing for each letter it's Value Type Unicode and HEX
		fmt.Printf("%v - %T - %#U - %#x\n", letter, letter, letter, letter)
	}
}

func GetOSInfo() {
	fmt.Printf("Using OS: %v\n", runtime.GOOS)
	fmt.Printf("Wiht Processor: %v\n", runtime.GOARCH)
}

func NumericSystems() {
	number := 2341
	// Printing Decimal / Binary / HEX of number
	fmt.Printf("Decimal: %d\nBinary: %b\nHex: %#x\n", number, number, number)
}

func ExampleConstants() {
	// Consts can have a Type but if we don't give it a Type it decuses the type when the variable is called
	// As you can see below we define x = 10 that normally would be a int/usigned but we give it to a float and it works
	// Cause the compiler deduces 10 as a float, but if we used a var int instead it wouldn't work

	const x = 10
	var y float64 = x
	fmt.Printf("%v - %T\n", y, y)
}

func IotasExample() {
	// Basically consts that are not initialized with values
	// Copy the same operation from the last Const initialized
	// And as Iota changes it value every time that it is called
	const (
		_ = iota
		KB = 1 << (iota * 10)
		MB
		GB
		TB
	)
	fmt.Printf("Iota KB: %v\nIota MB: %v\nIota GB: %v\nIota TB: %v\n", KB, MB, GB, TB)
}

func InitialFunctionsRunner() {
	StringTests()
	GoSwap()
	PrimitiveTypes()
	PritnPackageVariable()
	ZeroValuesForTypes()
	TypesForStrings()
	GetStringBytes("Jolestest")
	GetOSInfo()
	NumericSystems()
	ExampleConstants()
	IotasExample()
}