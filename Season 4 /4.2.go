package main

import (
	"fmt"
)

func countItems(lst []int) int {
	if len(lst) == 0 {
		return 0
	}
	return 1 + countItems(lst[1:])
}

func main() {
	items := []int{10, 20, 30, 40, 50}
	fmt.Println(countItems(items))
}
