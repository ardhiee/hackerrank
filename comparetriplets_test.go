package hackerrank

import (
	"reflect"
	"testing"
)

func TestCompareTriplets(t *testing.T) {

	t.Run("Size of array is 3 equal", func(t *testing.T) {
		a := []int{6, 6, 6}
		b := []int{2, 1, 3}

		got := compareTriplets(a, b)
		want := []int{3, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
