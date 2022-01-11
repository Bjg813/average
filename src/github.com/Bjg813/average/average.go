// Average calculates the average of several numbers.
package average

import (
	"fmt"
)

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	// Use len to find the length of the array.
	// Divide the sum of the array integers then divide by length of array to find the average.
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
