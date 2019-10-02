package hackerrank

// Sum will return total of arrays
func Sum(numbers []int) int {

	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
