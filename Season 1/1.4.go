package main

import (
	"fmt"
)

func linear_search(phoneBook [][2]string, name string) string {
	for _, contanct := range phoneBook {
		if contanct[0] == name {
			return contanct[1]
		}
	}
	return "Not Found"
}

func main() {
	phoneBook := [][2]string{
		{"Alice", "123456"},
		{"Bob", "234567"},
		{"Charlie", "345678"},
		{"David", "456789"},
		{"Eve", "567890"},
	}

	result := linear_search(phoneBook, "David")
	fmt.Println("Liner search:", result)
}
