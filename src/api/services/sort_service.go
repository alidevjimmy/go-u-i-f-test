package services

import "github.com/alidevjimmy/go-u-i-f-test/src/api/utils/sort"

func Sort(array []int) {
	if len(array) < 1000 {
		sort.BubbleSortDescRecursive(array)
		return
	}
	sort.Sort(array)
}
