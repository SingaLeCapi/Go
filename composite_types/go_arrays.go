package main
import "fmt"

func main(){
	nums := [3]int{1, 3, 5}
	modifyArray(&nums)	
	fmt.Println("New value of nums:", nums )
}

// Modifying the value of an array using a pointer
func modifyArray(arr *[3]int) {
	arr[0] = 11;
}

