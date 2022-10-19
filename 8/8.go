package main

import (
	"fmt"
	"strings"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	firstSymbols := make(map[string][]int)
	lastSymbols := make(map[string][]int)
	domains := make([]string, m, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&domains[i])
		runes := []rune(domains[i])
		firstSymbols[string(runes[0])] = append(firstSymbols[string(runes[0])], i)
		lastSymbols[string(runes[len(runes)-1])] = append(lastSymbols[string(runes[len(runes)-1])], i)
	}
	for i := 0; i < n; i++ {
		var start, end string
		var count int
		fmt.Scan(&start, &end)
		first := string([]rune(start)[0])
		last := string([]rune(end)[len([]rune(end))-1])
		if slice1, ok := firstSymbols[first]; ok {
			if slice2, ok := lastSymbols[last]; ok {
				commonSlice := findDuplicates(slice1, slice2)
				if len(commonSlice) > 0 {
					for _, number := range commonSlice {
						domainForCheck := domains[number]
						if strings.HasPrefix(domainForCheck, start) && strings.HasSuffix(domainForCheck, end) {
							count++
						}
					}
				}
			}
		}
		fmt.Println(count)
	}
}

func findDuplicates(slice1, slice2 []int) []int {
	commonSlice := append(slice1, slice2...)
	keys := make(map[int]bool)
	var list []int
	for _, value := range commonSlice {
		if _, ok := keys[value]; ok {
			list = append(list, value)
		} else {
			keys[value] = true
		}
	}
	return list
}
