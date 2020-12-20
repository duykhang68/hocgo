package main

import "fmt"

func main() {
	fmt.Println("Hello Golang")
	fmt.Println(Sum(3)) // o + 1 + 2
}
func Sum(number int) int {
	Sum := 0
	for i := 0; i < number; i++ {
		Sum += i

	}
	return Sum
}
