package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println("penjumlahan:", a*d+b*c, "/", b*d)
	fmt.Println("pengurangan:", a*d-b*c, "/", b*d)
	fmt.Println("perkalian:", a*c, "/", b*d)
	fmt.Println("pembagian:", a*d, "/", b*c)
}
