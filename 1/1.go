package main

import (
	"fmt"
)

func main() {
	input := [8]int{}
	for i := 0; i < len(input); i++ {
		fmt.Scan(&input[i])
	}
	minX, maxX := input[0], input[0]
	minY, maxY := input[0], input[0]

	for i := 0; i < len(input); i++ {
		for j := range input {
			//X
			if input[j] > maxX && j%2 == 0 {
				maxX = input[j]
			} else if input[j] < minX && j%2 == 0 {
				minX = input[j]
			}
			//Y
			if input[j] > maxY && j%2 != 0 {
				maxY = input[j]
			} else if input[j] < minY && j%2 != 0 {
				minY = input[j]
			}
		}
	}
	if (maxX - minX) > (maxY - minY) {
		fmt.Println((maxX - minX) * (maxX - minX))
	} else if (maxX - minX) < (maxY - minY) {
		fmt.Println((maxY - minY) * (maxY - minY))
	}
}
