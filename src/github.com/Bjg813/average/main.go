// Average calculates the average of several numbers using a variadic function.
package main

import "fmt"

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	fmt.Println(average(20, 10))
}
