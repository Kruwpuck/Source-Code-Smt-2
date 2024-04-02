package main

import "fmt"

func main() {
	var n, jumlah, masuk int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&masuk)
		jumlah += ((masuk % 100) / 10) + (masuk / 1000 % 10)
	}
	fmt.Print(jumlah)
}
