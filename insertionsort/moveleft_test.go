package insertionsort

import (
	"testing"

	"github.com/tel23a-inf/aufgaben-sortieren/testhelpers"
)

func TestMoveLef(t *testing.T) {
	list := []int{3, 2, 1}
	MoveLeft(list, 2)
	testhelpers.AssertListsEqual(t, []int{1, 3, 2}, list)
}