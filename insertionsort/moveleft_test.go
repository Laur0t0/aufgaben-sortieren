package insertionsort

import (
	"testing"

	"github.com/tel23a-inf/aufgaben-sortieren/testhelpers"
)

func TestMoveLef(t *testing.T) {
	list := []int{3, 4, 2}
	MoveLeft(list, 2)
	testhelpers.AssertListsEqual(t, []int{2, 3, 4}, list)
}
