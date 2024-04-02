package main

import "fmt"

func main() {
	var n, i float64
	i = 1
	for true {
		fmt.Scan(&n)
		if n != 0 {
			i *= (1.0 / n)
		} else {
			break
		}
	}
	fmt.Printf("%.5f", i)

}
