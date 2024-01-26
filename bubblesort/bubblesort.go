package bubblesort

// BubbleSort sorts the given list using the bubble sort algorithm.
func BubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		BubbleUp(list)
	}
}
