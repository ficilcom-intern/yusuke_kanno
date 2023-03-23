package main

import (
	"fmt"
)

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	sum := 0

	for i := 0; i <= 4; i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
