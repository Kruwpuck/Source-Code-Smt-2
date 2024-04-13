package main

import "fmt"

func main() {
	var a int
	var hasil bool = true
	fmt.Scan(&a)
	if a == 1 || a == 2 {
		fmt.Print(false)
	} else {
		for i := 2; i <= a-1; i++ {
			if a%i == 0 {
				hasil = false
			}
		}
		fmt.Print(!hasil)
	}

}
