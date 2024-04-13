package main

import (
	"fmt"
)

func main() {
	var a int
	var hasil bool = true
	fmt.Scan(&a)
	if a == 1 {
		fmt.Print(false)
	} else if a == 2 {
		fmt.Print(true)

	} else {
		for i := 3; i <= a-1; i++ {
			if a%i == 0 {
				hasil = false
			}
		}
		fmt.Print(hasil)
	}

}
