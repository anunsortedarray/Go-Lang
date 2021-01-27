// My first Go Program :)
package main

import "fmt"

func main() {
	fmt.Println("Hi, I'm new to Go.")
	foo()
	bar()
	fmt.Println("Bye.")
}

func foo() {
	fmt.Println("In Foo")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func bar() {
	fmt.Println("In Bar")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
