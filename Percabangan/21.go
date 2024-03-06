package main

import "fmt"

func main() {
	var a, b bool
	fmt.Scan(&a, &b)
	if a && b {
		fmt.Println("keluar jalan-jalan")
	} else {
		fmt.Println("diam di rumah saja")
	}
}
