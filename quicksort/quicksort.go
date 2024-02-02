package quicksort

// QuickSort sorts the given list using the quicksort algorithm.
func QuickSort(list []int) {
	if len(list) < 2 {
		return
	}

	pivot := list[0]
	left := []int{}
	right := []int{}

	for _, v := range list[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	QuickSort(left)
	QuickSort(right)

	copy(list, append(append(left, pivot), right...))
}
