package main

import "fmt"

func main() {
	var k1, k2, k10, masuk int
	fmt.Scan(&masuk)
	k10 = masuk / 10000
	k2 = (masuk % 10000) / 2000
	k1 = (masuk % 10000 % 2000) / 1000
	fmt.Println(k10, "lembar 10000")
	fmt.Println(k2, "lembar 2000")
	fmt.Println(k1, "lembar 1000")
}
