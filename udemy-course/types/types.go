package main

import "fmt"

var z string = "Hello"

func main() {
	x := "Hello"
	fmt.Println("Value of x = ", x)
	fmt.Printf("Type of x = %T\n", x)

	var y float64 = 3.14
	fmt.Println("Value of y = ", y)
	fmt.Printf("Type of y = %T\n", y)

	// As go is a static typed language, the following is not valid
	// z = 43
	// fmt.Println("Value of z = ", z)
}
