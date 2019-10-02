package hackerrank

import "testing"

func TestSolveMeFirst(t *testing.T) {

	var want uint32
	got := solveMeFirst(2, 3)
	want = 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
