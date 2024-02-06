// Program Kelipatan 15
package main

import "fmt"

func main() {
	var X int
	var stop, hasil bool
	stop = false
	fmt.Scan(&X)
	for !stop {
		X -= 15
		if X == 0 {
			hasil = true
		}
		if X <= 0 {
			stop = true
		}
	}
	fmt.Println(hasil)
}
