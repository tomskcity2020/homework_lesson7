package main

import (
	"fmt"
)

func main() {
	xy := 8
	result := ""
	value := " "

	for y := 0; y < xy; y++ {
		if y%2 == 0 {
			for x := 0; x < xy; x++ {
				if x%2 == 0 {
					value = " "
				} else {
					value = "#"
				}
				result += value

			}
		} else {
			for x := 0; x < xy; x++ {
				if x%2 != 0 {
					value = " "
				} else {
					value = "#"
				}
				result += value
			}
		}

		result += "\n"

	}

	fmt.Println(result)
}
