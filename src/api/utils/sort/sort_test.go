package sort

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	var array = []int{43, 123, 81, 2, 6, 2, 9}

	Sort(array)

	if !sort.IntsAreSorted(array) {
		unSortedArray := array
		sort.Ints(array)
		t.Errorf("expected array is %v but it is %v\n", array, unSortedArray)
	}
}

func BenchmarkSort(b *testing.B) {
	var array = []int{43, 123, 81, 2, 6, 2, 9}
	for i := 0; i < b.N; i++ {
		Sort(array)
	}
}
