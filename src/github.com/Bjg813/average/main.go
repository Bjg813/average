// Average calculates the average of several numbers using a variadic function and a slice.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	// Arguments variable will hold the numbers entered from the terminal.
	// Numbers slice will eventually hold the numbers we're averaging entered from the terminal.
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		// Append the converted number to the numbers slice.
		numbers = append(numbers, number)
	}
	// Pass the slice to the variadic function and make turn the numbers slice into a variadic argument using (...) at end of slice numbers name.
	fmt.Printf("Average: %0.2f\n", average(numbers...))
	// Must use "go build main.go" in order to use this program in the terminal.
	// Use ./main then write the numbers you want to average.
	// Example: ./main 15 30 65
	// 36.67
}
