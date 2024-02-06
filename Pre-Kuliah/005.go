// Program Sport Club
package main

import "fmt"

func main() {
	var S, hasil int
	fmt.Scan(&S)
	hasil = S / 7
	if S%7 != 0 {
		hasil += 1
	}
	fmt.Println(hasil)
}
