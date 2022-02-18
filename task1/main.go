package main

import "fmt"

func main() {

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if isEven(number) {
			fmt.Println(number, " is even", number)
		} else {
			fmt.Println(number, " is odd", number)
		}
	}
}

func isEven(number int) bool {
	return number%2 == 0
}
