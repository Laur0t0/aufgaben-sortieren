package mergesort

// Merge merges the two sorted sublists into a single sorted list.
func Merge(list1, list2 []int) []int {
	result := []int{}
	i := 0
	j := 0
	for i < (len(list1)) && i < (len(list2)) {
		if list1[i] < list2[j] {
			result = append(result, list1[i])
			i++
		} else {
			result = append(result, list2[j])
			j++
		}
	}
	if i > len(list1) && j > len(list2) {
		result = append(result, list1[i])
	} else {
		result = append(result, list2[j])
	}
	return result
}
