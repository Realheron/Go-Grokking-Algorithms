
package main

import (
	"fmt"
)

func findMaxRecursive(lst []int) int {
	if len(lst) == 1 {
		return lst[0]
	}
	maxRest := findMaxRecursive(lst[1:])
	if lst[0] > maxRest {
		return lst[0]
	}
	return maxRest
}

func main() {
	numbers := []int{4, 10, 2, 33, 5}
	fmt.Println(findMaxRecursive(numbers))
}
