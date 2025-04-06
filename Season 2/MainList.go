package main

import (
	"fmt"
)

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	newArr := []int{}

	copyArr := make([]int, len(arr))
	copy(copyArr, arr)

	for len(copyArr) > 0 {
		smallestIndex := findSmallest(copyArr)
		newArr = append(newArr, copyArr[smallestIndex])

		copyArr = append(copyArr[:smallestIndex], copyArr[smallestIndex+1:]...)
	}

	return newArr
}

func main() {
	result := selectionSort([]int{5, 3, 6, 2, 10})
	fmt.Println(result)
}

