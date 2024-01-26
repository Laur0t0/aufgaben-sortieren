package selectionsort

// SmallestPos returns the index of the smallest element in the given list.
func SmallestPos(list []int) int {
	smallest := 0
	for i, value := range list {
		if value < list[smallest] {
			smallest = i
		}
	}
	return smallest
}
