package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n, q int
	surnames := make(map[string]int) //KEY - prefix, VALUE - index in original slice
	сards := make(map[byte][]string) //KEY - starter unicode character, VALUE - words starting with this character
	fmt.Scan(&n, &q)
	for i := 0; i < n; i++ {
		var input string
		fmt.Scan(&input)
		surnames[input] = i + 1
		сards[input[0]] = append(сards[input[0]], input)
	}
	for i := range сards {
		sort.Strings(сards[i])
	}
	//---------------------------------
	answer := []int{}
	var prefixExist bool
	for i := 0; i < q; i++ {
		var number, counter int
		var prefix string
		fmt.Scan(&number, &prefix)
		for _, value := range сards[prefix[0]] {
			if strings.HasPrefix(value, prefix) {
				counter++
				prefixExist = true
				if counter == number {
					answer = append(answer, surnames[value])
				}
			}
		}
	}
	//---------------------------------
	if prefixExist {
		for _, v := range answer {
			fmt.Println(v)
		}
	} else {
		fmt.Println(-1)
	}
}
