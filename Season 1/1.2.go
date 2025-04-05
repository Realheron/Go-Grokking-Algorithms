package main

import (
	"fmt"
)

func binarySearch(list []string, item string) (int, int) {
	low := 0
	high := len(list) - 1
	steps := 0

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		steps++

		if guess == item {
			return mid, steps
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, steps
}

func main() {

	NameList := make([]string, 256)
	for i := 0; i < 256; i++ {
		NameList[i] = fmt.Sprintf("Name_%d", i)
	}

	index, steps := binarySearch(NameList, "Name_75")
	fmt.Printf("index: %d, steps: %d\n", index, steps)
}
