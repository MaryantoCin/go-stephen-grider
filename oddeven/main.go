package main

import "fmt"

func main() {
	numbers := []int{}

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for _, number := range numbers {
		status := ""
		if number%2 == 0 {
			status = "even"
		} else {
			status = "odd"
		}
		fmt.Printf("%d is %v\n", number, status)
	}
}
