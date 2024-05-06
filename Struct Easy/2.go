package main

import "fmt"

func main() {
	type pemain struct {
		nama        string
		gol, assist float64
	}
	var a, b, c pemain
	fmt.Scan(&a.nama, &a.gol, &a.assist)
	fmt.Scan(&b.nama, &b.gol, &b.assist)
	fmt.Scan(&c.nama, &c.gol, &c.assist)
	fmt.Println((a.gol + b.gol + c.gol))
	fmt.Println((a.assist + b.assist + c.assist))
}
