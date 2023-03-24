package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello World!")
}

func main() {

	func(greeting string) {
		fmt.Println(greeting)
	}("Good evening")
}
