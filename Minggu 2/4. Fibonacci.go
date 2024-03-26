package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibo(n))
}
func fibo(n int) int {
	var i1, i0, hasil int
	var jumlah int
	i0 = 0
	i1 = 1
	for i := 2; i <= n; i++ {

		hasil = i1 + i0
		i0 = i1
		i1 = hasil
		jumlah += i0
	}
	return jumlah
}
