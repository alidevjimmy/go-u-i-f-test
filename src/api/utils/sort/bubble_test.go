package sort

import (
	"testing"
	"time"
)

func TestBubbleSortDescRegularSorted(t *testing.T) {
	array := []int{3, 7, 35, 8, 2, 1, 8, 9}

	timeOutChan := make(chan bool, 1)
	defer close(timeOutChan)

	go func() {
		BubbleSortDescRegular(array)
		timeOutChan <- false
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		timeOutChan <- true
	}()

	if <-timeOutChan {
		t.Fatal("sort timeout")
		return
	}

	if array[0] != 35 {
		t.Error("the first element of desc sorted array must 35")
	}
	if array[len(array)-1] != 1 {
		t.Error("the last element of desc sorted array must 1")
	}

}

func TestBubbleSortAscRegularSorted(t *testing.T) {
	array := []int{3, 7, 35, 8, 2, 1, 8, 9}
	BubbleSortAscRegular(array)
	if array[0] != 1 {
		t.Error("the first element of desc sorted array must 1")
	}
	if array[len(array)-1] != 35 {
		t.Error("the last element of desc sorted array must 35")
	}
}

func TestBubbleSortDescRecursiveSorted(t *testing.T) {
	array := []int{3, 7, 35, 8, 2, 1, 8, 9}
	BubbleSortDescRecursive(array)
	// assert.EqualValues(1, array[len(array)-1])

	if array[len(array)-1] != 1 {
		t.Error("the last element of desc sorted array must 1")
	}
}

func BenchmarkBubbleSortAscRegular(b *testing.B) {
	var array = []int{43, 123, 81, 2, 6, 2, 9}
	for i := 0; i < b.N; i++ {
		BubbleSortAscRegular(array)
	}
}
