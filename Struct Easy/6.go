package main

import "fmt"

func main() {
	type tim struct {
		nama           string
		point, selisih float64
	}
	var a, b, c tim
	fmt.Scan(&a.nama, &a.point, &a.selisih)
	fmt.Scan(&b.nama, &b.point, &b.selisih)
	fmt.Scan(&c.nama, &c.point, &c.selisih)
	fmt.Println((a.selisih + b.selisih + c.selisih) / 3)

}
