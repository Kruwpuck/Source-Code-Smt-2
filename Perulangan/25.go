package main

import "fmt"

func main() {
	var n, aset, prib, masuk int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&masuk)
		prib += masuk / 3
		aset += masuk % 3
	}
	fmt.Println(prib, aset)
}
