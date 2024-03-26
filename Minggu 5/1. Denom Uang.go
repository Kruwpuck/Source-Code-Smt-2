package main

import "fmt"

func main() {
	var uang, k1, k10, k2 int
	fmt.Scan(&uang)
	pecahUang(uang, &k10, &k2, &k1)
	fmt.Println(k10, k2, k1)
}
func pecahUang(uang int, k10, k2, k1 *int) {
	*k10 = uang / 10000
	*k2 = (uang % 10000) / 2000
	*k1 = (uang % 10000 % 2000)
}
