package hackerrank

import "testing"

func TestSimpleArraySum(t *testing.T) {

	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
