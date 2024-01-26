package insertionsort

// InsertionSort sorts the given list using the insertion sort algorithm.
func InsertionSort(list []int) {
	for i := 0; i < len(list); i++ {
		MoveLeft(list, len(list))
	}
}
