package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a < 100 {
		fmt.Print(a * 12)
	} else if a < 200 {
		fmt.Print(a*12*20/100 + a*12)
	} else if a >= 200 && a < 400 {
		fmt.Print(a*15*20/100 + a*15)
	} else if a >= 400 && a < 600 {
		fmt.Print((a * 18 * 20 / 100) + a*18)
	} else if a > 600 {
		fmt.Print((a * 20 * 20 / 100) + a*20)
	}
}
