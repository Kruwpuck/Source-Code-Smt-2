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
	if a.gol < b.gol && a.gol < c.gol {
		fmt.Println(a.nama)
	} else if b.gol < a.gol && b.gol < c.gol {
		fmt.Println(b.nama)
	} else if c.gol < a.gol && c.gol < b.gol {
		fmt.Println(c.nama)
	}
	if a.assist < b.assist && a.assist < c.assist {
		fmt.Println(a.nama)
	} else if b.assist < a.assist && b.assist < c.assist {
		fmt.Println(b.nama)
	} else if c.assist < a.assist && c.assist < b.assist {
		fmt.Println(c.nama)
	}
}
