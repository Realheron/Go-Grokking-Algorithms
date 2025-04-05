package main

import (
	"fmt"
)

func lookForKey(box []interface{}) {
	for _, item := range box {
		switch val := item.(type) {
		case string:
			if val == "key" {
				fmt.Println("Found the key!")
			}
		case []interface{}:
			lookForKey(val)
		}
	}
}

func main() {

	myBox := []interface{}{
		"item1",
		[]interface{}{
			"item2",
			[]interface{}{
				"key",
			},
		},
		"item3",
	}

	lookForKey(myBox)
}
