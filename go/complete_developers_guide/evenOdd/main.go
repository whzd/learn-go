package main

import "fmt"

func main() {
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range n {
		if value%2 != 0 {
			fmt.Printf("%v is odd\n", value)
		} else {
			fmt.Printf("%v is even\n", value)
		}
	}
}
