package main

import (
	"fmt"
)
func sumRecursive(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sumRecursive(arr[1:])
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sumRecursive(numbers))

}
