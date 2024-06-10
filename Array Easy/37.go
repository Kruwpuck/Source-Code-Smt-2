package main

import "fmt"

func main() {
	var duid, k10, k5, k1 int
	fmt.Scan(&duid)
	k10 = duid / 10000
	k5 = (duid % 10000) / 5000
	k1 = (duid % 10000 % 5000) / 1000
	fmt.Println(k10, k5, k1)
}
