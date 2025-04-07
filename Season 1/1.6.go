package main

import (
	"fmt"
	"strings"
)

func main() {
	phoneBook := [][2]string{
		{"Alice", "135659"},
		{"Bob", "4578213"},
		{"Charlie", "468769"},
		{"David", "4651358"},
		{"Eve", "1254687"},
		{"Heron", "589746"},
	}

	var aNames []string

	for _, contact := range phoneBook {
		name := contact[0]
		number := contact[1]
		if strings.HasPrefix(name, "A") {
			aNames = append(aNames, number)
		}
	}
	fmt.Println("Numbers of names starting with 'A':", aNames)
}
