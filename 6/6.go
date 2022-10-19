package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	chains := make(map[int][]int)
	var lengths []int
	for i := 0; i < n; i++ {
		var start, end int
		fmt.Scan(&start, &end)
		chains[start] = append(chains[start], end)
	}
	for key := range chains {
		sort.Ints(chains[key])
	}
	count(0, 0, chains, &lengths)
	sort.Ints(lengths)
	fmt.Println(lengths[len(lengths)-1])
}

func count(actual int, counter int, chains map[int][]int, lengths *[]int) {
	if _, ok := chains[actual]; !ok {
		*lengths = append(*lengths, counter)
		return
	}
	counter++
	for _, value := range chains[actual] {
		if value != actual {
			count(value, counter, chains, lengths)
		} else {
			counter++
		}
	}
}
