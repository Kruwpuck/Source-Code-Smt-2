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
	if a.point > b.point && a.point > c.point {
		fmt.Println(a.nama)
	} else if b.point > a.point && b.point > c.point {
		fmt.Println(b.nama)
	} else {
		fmt.Println(c.nama)
	}
	if a.selisih > b.selisih && a.selisih > c.selisih {
		fmt.Println(a.nama)
	} else if b.selisih > a.selisih && b.selisih > c.selisih {
		fmt.Println(b.nama)
	} else {
		fmt.Println(c.nama)
	}
}
