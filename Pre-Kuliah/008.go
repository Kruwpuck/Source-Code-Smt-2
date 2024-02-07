package main

import "fmt"

func main() {
	var n, harga int
	var diskon, total float32
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&harga, &diskon)
		total += (float32(harga) * (100 - (diskon)) / 100)
	}
	fmt.Println(int(total))
}
