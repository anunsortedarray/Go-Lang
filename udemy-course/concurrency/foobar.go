package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Entering main go routine")
	go foo(10)
	go bar(10)
	wg.Add(2)
	fmt.Println("Exiting main go routine")
	wg.Wait()
}

func foo(count int) {
	for i := 0; i < count; i++ {
		fmt.Println("Foo: ", i)
	}
	wg.Done()
}

func bar(count int) {
	for i := 0; i < count; i++ {
		fmt.Println("Bar: ", i)
	}
	wg.Done()
}
