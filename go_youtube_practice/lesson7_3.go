package main

import (
	"fmt"
	"reflect"
)

func main() {
	var string_a string = "Hello World!"
	string_b := "Hello World!"
	fmt.Println(string_a)
	fmt.Println(reflect.TypeOf(string_a))
	fmt.Println(string_b)
	fmt.Println(reflect.TypeOf(string_b))
}
