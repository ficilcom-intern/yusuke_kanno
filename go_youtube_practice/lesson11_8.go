package main

import (
	"fmt"
)

func main() {
	i := 0

	for i <= 4 {
		fmt.Println(i)
		if i == 4 {
			break
		}
		i++
	}
}
