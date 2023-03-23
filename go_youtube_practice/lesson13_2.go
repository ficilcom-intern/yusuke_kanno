package main

import (
	"fmt"
)

type Student struct{
	name string
	math, english float64
}

func main() {
	s := Student{name: "sato", math: 80, english: 70}

	fmt.Println(s)
}
