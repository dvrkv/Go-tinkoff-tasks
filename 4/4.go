package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	items := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	helper(items, scanner)
}

func helper(lastItems map[string]string, scanner *bufio.Scanner) {
	items := make(map[string]string)
	for key, value := range lastItems {
		items[key] = value
	}
	for scanner.Scan() {
		item := scanner.Text()
		if len(item) == 0 {
			break
		}
		if item == "{" {
			helper(items, scanner)
		}
		if item == "}" {
			break
		}
		temp := strings.Split(item, "=")
		if len(temp) == 2 {
			if _, er := strconv.Atoi(temp[1]); er == nil {
				items[temp[0]] = temp[1]
			} else {
				if _, ok := items[temp[1]]; ok {
					items[temp[0]] = items[temp[1]]
				} else if _, ok := lastItems[temp[1]]; ok {
					items[temp[0]] = lastItems[temp[1]]
				} else {
					items[temp[0]] = "0"
				}
				fmt.Println(items[temp[0]])
			}
		}
	}
}
