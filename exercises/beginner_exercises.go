package exercises

import "fmt"

func BeginnerExercisesRunner() {
	fmt.Println("=========Performing Beginner Exercises=========")
	exercise_one()
	fmt.Println()
	exercise_two()
	fmt.Println()
	exercise_three()
	fmt.Println()
	exercise_four()
	fmt.Println()
	exercise_five()
}

func exercise_one() {
	println("==========EXERCISE ONE==============")

	x, y, z :=  42, "James Bond", true
	fmt.Println("Value of x: ", x)
	fmt.Println("Value of y: ", y)
	fmt.Println("Value of z: ", z)
	fmt.Printf("All togheter: %v, %v, %v", x, y, z)
}

// Will be used later
var x int
var y string
var z bool

func exercise_two() {
	println("==========EXERCISE TWO==============")

	fmt.Printf(
		"Values and types of variables:\nX: %v, %T\nY: %v, %T\nZ: %v, %T", 
		x, x, 
		y, y, 
		z, z,
	)
}

func exercise_three() {
	println("==========EXERCISE THREE==============")

	x = 42
	y = "James Bond"
	z = true
	formatted_string := fmt.Sprintf("Name: %v, Age: %v, Secret Agent: %v", y, x, z)
	fmt.Print(formatted_string)
}

type good_int int
var another_x good_int

func exercise_four() {
	println("==========EXERCISE FOUR==============")

	fmt.Printf("X before assignment with type: %v, %T\n", another_x, another_x)
	another_x = 42
	fmt.Printf("X after assignment: %v", another_x)
}

var another_y int

func exercise_five() {
	println("==========EXERCISE FIVE==============")
	fmt.Printf("X before assignment with type: %v, %T\n", another_x, another_x)
	another_x = 42
	fmt.Printf("X after assignment: %v\n", another_x)

	fmt.Printf("Y before assignment with type: %v, %T\n", another_y, another_y)
	// Converting custom type to subtype
	another_y = int(another_x)
	fmt.Printf("Y after assignment with type: %v, %T", another_y, another_y)
}