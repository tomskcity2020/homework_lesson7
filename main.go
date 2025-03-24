package main

import (
	"fmt"
	"math/rand"
)

func main() {
	xy := 4
	result := ""
	array := []string{" ", "#"}

	for y := 0; y < xy; y++ {

		for x := 0; x < xy; x++ {
			key := rand.Intn(2)
			get_value := array[key]
			result += get_value
		}

		result += "\n"

	}

	fmt.Println(result)
}
