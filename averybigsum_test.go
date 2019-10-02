package hackerrank

import "testing"

func TestAveryBigSum(t *testing.T) {

	ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	var want int64

	got := aVeryBigSum(ar)
	want = 5000000015

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
