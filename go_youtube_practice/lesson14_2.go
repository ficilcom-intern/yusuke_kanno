package main

import (
	"fmt"
)

type Student struct{
	name string
	math, english float64
}

func (s Student) avg(math, english float64) {
	fmt.Println((s.math + s.english) / 2)
}

func main() {
	a001 := Student{"sato", 80, 70}
	a001.avg(80, 70)
}
