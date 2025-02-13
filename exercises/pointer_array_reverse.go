package exercises

import "fmt"

func ReverseArrayUsingPointers(arr [20]int ) [20]int {
	left, right := 0, len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	fmt.Println("Reversed array: ", arr)
	return arr
}