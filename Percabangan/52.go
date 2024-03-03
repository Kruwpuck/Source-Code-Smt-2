package main

import "fmt"

func main() {
	var asli, lahir, imigrasi, mati, remigrasi int
	fmt.Scan(&asli, &lahir, &imigrasi, &mati, &remigrasi)
	fmt.Println(asli + lahir + imigrasi - mati - remigrasi)
}
