package hackerrank

func aVeryBigSum(ar []int64) int64 {

	var sum int64

	for _, number := range ar {
		sum += number
	}

	return sum
}
