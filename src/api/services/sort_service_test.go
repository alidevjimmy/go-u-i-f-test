package services

import (
	"testing"

	"github.com/alidevjimmy/go-u-i-f-test/src/api/utils/sort"
)

func TestSort(t *testing.T) {
	array := []int{3, 7, 35, 8, 2, 1, 8, 9}
	Sort(array)
	if array[0] != 35 {
		t.Error("the first element of desc sorted array must 35")
	}
	if array[len(array)-1] != 1 {
		t.Error("the last element of desc sorted array must 1")
	}
}

func getArray(l int) []int {
	array := make([]int, l)
	for i := 1; i <= l; i++ {
		array[i-1] = i
	}
	return array
}
func TestSortLessAndEqualThan1000Desc(t *testing.T) {
	array := getArray(1000)
	sort.BubbleSortDescRegular(array)
	if array[0] != 1000 {
		t.Error("expected 1000!")
	}

	if array[len(array)-1] != 1 {
		t.Error("expected 1!")
	}
}

func TestSortMoreThan1000(t *testing.T) {
	array := getArray(10000)
	sort.BubbleSortDescRegular(array)
	if array[0] != 10000 {
		t.Error("expected 10000!")
	}

	if array[len(array)-1] != 1 {
		t.Error("expected 1!")
	}
}

