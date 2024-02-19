package insertionsort

// MoveLeft moves the element at the given position to the left until it is in the correct position.
func MoveLeft(list []int, pos int) {

	// if len(list) > 1 {
	// 	for i := len(list) - 2; i != -1; i-- {
	// 		for list[i] > list[i+1] {
	// 			list[i], list[i+1] = list[i+1], list[i]
	// 		}
	// 	}
	// }

	if len(list) > 1 {
		for i := pos; i > 0; i-- {							//Starte an der gegebenen Stelle pos
			if list[i] < list[i-1] {						//wenn 
				list[i], list[i-1] = list[i-1], list[i]
			}
		}
	}
}
