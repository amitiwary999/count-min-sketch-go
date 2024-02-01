package main

import (
	cms "countminsketch/count-min"
	"fmt"
)

func main() {
	minsketch := cms.NewCountMin(1000)

	for i := 0; i < 100000; i++ {
		key := "min-sketch"
		modifiedKey := ""
		if i%50 == 0 {
			modifiedKey = fmt.Sprintf("%v-%v", key, 50)
		} else if i%10 == 0 {
			modifiedKey = fmt.Sprintf("%v-%v", key, 10)
		} else {
			modifiedKey = fmt.Sprintf("%v-%v", key, i)
		}
		minsketch.SetKeyCount(modifiedKey)

	}

	key := fmt.Sprintf("min-sketch-10")
	count := minsketch.GetKeyCount(key)
	fmt.Printf("count of the min sketch %v \n", count)
}
