package exercises

import "fmt"

func BasicExercisesRunner() {
	fmt.Println("=========Performing Basic Exercises=========")
	exercise_six()
	fmt.Println()
	exercise_seven(10, 29)
	fmt.Println()
	exercise_eight()
	fmt.Println()
	exercise_nine()
	fmt.Println()
	exercise_ten()
	fmt.Println()
	exercise_eleven()
}

func exercise_six() {
	println("==========EXERCISE SIX==============")
	x := 192
	fmt.Printf("Decimal: %d\nBinary: %b\nHEX: %#x of %v", x, x, x, x)
}

func exercise_seven(number_one int, number_two int) {
	println("==========EXERCISE SEVEN==============")
	x, y, z := (number_one > number_two), (number_one < number_two), (number_two == number_one)
	fmt.Printf("%v > %v: %v\n", number_one, number_two, x)
	fmt.Printf("%v < %v: %v\n", number_one, number_two, y)
	fmt.Printf("%v == %v: %v\n", number_one, number_two, z)
}

func exercise_eight() {
	println("==========EXERCISE EIGHT==============")
	const (
		x int = 190
		y = iota
		z
	)
	fmt.Printf("%v - %T\n%v - %T\n%v - %T", x, x, y, y, z, z)
}

func exercise_nine() {
	println("==========EXERCISE NINE==============")
	x := 13
	fmt.Printf("%d, %b, %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d, %b, %#x", y, y, y)
}

func exercise_ten() {
	println("==========EXERCISE TEN==============")
	raw_string := `
		Jolestest
			Esteve
				Aqui`
	fmt.Println(raw_string)
}

func exercise_eleven() {
	println("==========EXERCISE ELEVEN==============")
	const (
		_ = iota
		X
		Y
		Z
		A
	)
	fmt.Println(X, Y, Z, A)
}