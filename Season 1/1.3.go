package main

import (
	"fmt"
)

func binarySearch(phoneBook [][2]string, name string) string {
	low := 0
	high := len(phoneBook) - 1

	for low <= high {
		mid := (low + high) / 2
		if phoneBook[mid][0] == name {
			return phoneBook[mid][1]
		} else if phoneBook[mid][0] > name {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return "Not found"
}

func main() {
	phoneBook := [][2]string{
		{"Alice", "123456"},
		{"Bob", "234567"},
		{"Charlie", "345678"},
		{"David", "456789"},
		{"Eve", "567890"},
	}

	result := binarySearch(phoneBook, "Alice")
	fmt.Println("Binary search:", result)

}
