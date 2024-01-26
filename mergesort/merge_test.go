package mergesort

import (
	"testing"

	"github.com/tel23a-inf/aufgaben-sortieren/testhelpers"
)

func TestMerge(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9}
	list2 := []int{2, 4, 6, 8, 10}

	result := Merge(list1, list2)
	testhelpers.AssertListsEqual(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, result)
}
