package exercises

func adjacentDuplicatesRemoval(arr []int) []int {
	if len(arr) == 0 {
		return  arr
	}
	write := 1
	for read := 1; read < len(arr); read++ {
		if arr[read - 1] != arr[read] {
			arr[write] = arr[read]
			write++
		}
	}
	return arr[:write]
}