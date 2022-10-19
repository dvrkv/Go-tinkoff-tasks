//done
package main

import "fmt"

func main() {
	//длина среза
	var n int
	fmt.Scan(&n)
	var minPositive int = 100
	var mostNegative int = 0
	var totalIncome int = 0

	a := make([]int, n) //создаем динамический срез
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i]) //считываем элементы среза
		if i%2 == 0 {
			totalIncome += a[i]
			if a[i] < minPositive {
				minPositive = a[i]
			}
		} else {
			totalIncome -= a[i]
			if a[i] > mostNegative {
				mostNegative = a[i]
			}
		}
	}
	fmt.Println(totalIncome + (mostNegative-minPositive)*2)
}
