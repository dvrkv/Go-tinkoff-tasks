package main

import (
	"fmt"
	"sort"
)

func main() {
	var input string
	var q int
	fmt.Scan(&input, &q)
	alphabet := make(map[string][]int)
	for index, value := range []rune(input) {
		key := string(value)
		alphabet[key] = append(alphabet[key], index+1)
	}
	letters := make([]string, 0, len(alphabet))
	for letter := range alphabet {
		letters = append(letters, letter)
	}
	sort.Strings(letters)
	var start, end, counter, pointer int
	for i := 0; i < q; i++ {
		fmt.Scan(&start, &end)
		useful := end - start + 1
		counter -= start - 1
		for _, key := range letters {
			isChanged := false
			for _, val := range alphabet[key] {
				if val < start || val > end {
					continue
				}
				if val < pointer {
					isChanged = true
				}
				pointer = val
			}
			if isChanged {
				counter = (counter/useful+1)*useful + pointer - start
			} else {
				counter = (counter/useful)*useful + pointer - start
			}
		}
	}
	fmt.Println(counter)
}
