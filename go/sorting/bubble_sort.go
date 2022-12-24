package sorting

import "fmt"

func bubbleSort() {
	var arr = []int{1, 39, 2, 9, 7, 54, 11}
	var isDone = false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(arr)-1 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isDone = false
			}
			i++
		}
	}

	fmt.Println(arr)
}
