package main

import (
	"fmt"
)

func hash_search(phoneDict map[string]string, name string) string {
	if val, exists := phoneDict[name]; exists {
		return val
	}
	return "Not found"
}

func main() {

	phoneDict := map[string]string{
		"Alice":   "123456",
		"Bob":     "234567",
		"Charlie": "345678",
		"David":   "456789",
		"Eve":     "567890",
	}

	result := hash_search(phoneDict, "David")
	fmt.Println("Hash search:", result)
}
