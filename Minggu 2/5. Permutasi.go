package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x >= y {
		fmt.Println(faktor(x), faktor(y), mutasi(x, y))
	} else if x <= y {
		fmt.Println(faktor(x), faktor(y), mutasi(y, x))
	}

}
func faktor(n int) int {
	var jumlah int
	jumlah = 1
	for i := 1; i <= n; i++ {
		jumlah *= i
	}
	return jumlah
}
func mutasi(x, y int) int {
	var mut int
	mut = faktor(x) / faktor(x-y)
	return mut
}
