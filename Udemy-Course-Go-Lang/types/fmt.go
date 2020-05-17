package main

import "fmt"

var x int = 12

func main() {
	fmt.Println("Value of x = ", x)
	fmt.Printf("Type of x = %T\n", x)
	fmt.Printf("Binary of x = %b\n", x)
	fmt.Printf("Hex of x = %#x\n", x)
	fmt.Printf("Octal of x = %O\n", x)
	fmt.Printf("Type of x = %U\n", x)

	s := fmt.Sprintf("Binary of x = %b\n", x)
	fmt.Println("S = ", s)
}
