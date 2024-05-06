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
	if a.gol+a.assist > b.gol+b.assist && a.gol+a.assist > c.gol+c.assist {
		fmt.Println(a.nama)
	} else if b.gol+b.assist > a.gol+a.assist && b.gol+b.assist > c.gol+c.assist {
		fmt.Println(b.nama)
	} else {
		fmt.Println(c.nama)
	}

}
