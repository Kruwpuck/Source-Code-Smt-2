package main

import "fmt"

func main() {
	var minum, makan, bayar int
	var tips bool
	fmt.Scan(&makan, &minum, &tips)
	bayar += makan + minum
	if tips {
		bayar += 5000
	}
	fmt.Println(bayar)
}
