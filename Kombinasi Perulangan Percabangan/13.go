package main

import "fmt"

func main() {
	var rata, total, input float64
	for i := 0; i < 4; i++ {
		fmt.Scan(&input)
		total += input
	}
	rata = total / 4
	if rata >= 3.50 {
		fmt.Print("bagus")
	} else if rata <= 1.50 {
		fmt.Print("kurang")
	} else if rata > 1.50 && rata < 3.50 {
		fmt.Print("sedang")
	}
}
