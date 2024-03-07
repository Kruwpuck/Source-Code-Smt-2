package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x <= 5 {
		fmt.Print(x * 1000)
	} else if x > 5 && x <= 10 {
		fmt.Print(x * 2000)
	} else {
		fmt.Print("cabut keanggotaan")
	}
}
