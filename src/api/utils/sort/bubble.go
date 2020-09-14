package sort

// BubbleSortDescRegular sorting array by Desc and regular
func BubbleSortDescRegular(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

// BubbleSortAscRegular sorting array by Asc and regular
func BubbleSortAscRegular(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

// BubbleSortDescRecursive sorting array by Desc and recursive
func BubbleSortDescRecursive(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		if array[i] < array[i+1] {
			array[i], array[i+1] = array[i+1], array[i]
			return BubbleSortDescRecursive(array)
		}
	}

	return array
}
