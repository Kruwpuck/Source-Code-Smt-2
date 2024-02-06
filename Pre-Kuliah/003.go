// Program Berang-Berang
package main

import "fmt"

func main() {
	var N, harga, total int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		fmt.Scan(&harga)
		total += (harga * 90 / 100)
	}
	fmt.Println(total)
}
