package main

import "fmt"

func main() {
	var p, l, t, b int
	var hasil bool
	fmt.Scan(&p, &l, &t, &b)
	hasil = p*l*t <= 1000000 && b <= 30000
	fmt.Println(hasil)
}
