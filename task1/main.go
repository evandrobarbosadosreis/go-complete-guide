package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if isEven(number) {
			fmt.Printf("%d is even\n", number)
		} else {
			fmt.Printf("%d is odd\n", number)
		}
	}
}

func isEven(number int) bool {
	return number%2 == 0
}
