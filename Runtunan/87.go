package main

import "fmt"

func main() {
	var x, d1, d3, d4 int
	var disc, cashback, voucher bool
	fmt.Scan(&x)
	d1 = x / 1000
	d3 = (x / 10) % 10
	d4 = x % 10
	disc = d3%2 == 0
	cashback = d4 != 0 && (d1+d3)%d4 == 0
	voucher = (disc == true && cashback == false) || (disc == false && cashback == true)
	fmt.Println(disc, cashback, voucher)
}
