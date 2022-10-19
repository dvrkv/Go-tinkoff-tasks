package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func findMax(a []int) {
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)
}

func main() {
	//--------------------------------
	var l int
	elements := make([][]string, l)
	elementMap := map[string]int{}
	//--------------------------------
	fmt.Scan(&l)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i <= l; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var s []string
		s = append(s, input)
		elements = append(elements, s)
	}
	//--------------------------------
	var a []string
	var b []string
	for i := range elements {
		for j := range elements[i] {
			a = strings.Split(elements[i][j], " ")
			sort.Strings(a)

			str1 := strings.Join(a, " ")
			b = append(b, str1)
			elementMap[str1]++
		}
	}
	//--------------------------------
	delete(elementMap, " ")
	delete(elementMap, "")
	var arr []int
	for _, v := range elementMap {
		arr = append(arr, v)
	}
	findMax(arr)
}
