package bubblesort

// BubbleUp implements a single iteration of the bubble sort algorithm.
// It iterates over the list and compares each element with the next one.
// If the element is larger than the next one, they are swapped.
// The function returns true if at least one swap was performed.
func BubbleUp(list []int) bool {
	swapped := false
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
			swapped = true
		}
	}
	return swapped
}