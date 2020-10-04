package main

import "fmt"

func main() {

	fmt.Println("Hi This is Sushil learning GO and making history")

	foo()

	fmt.Println("Doing Something More")

	for i := 0; i < 100; i++ {

		if i%2 == 0 {

			fmt.Println(i)
		}
	}

	bar()
}

func foo() {

	fmt.Println("I am in FOO")
}

func bar() {

	fmt.Println("and then we exited")
}
