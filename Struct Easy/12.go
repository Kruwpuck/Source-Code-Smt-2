package main

import (
	"fmt"
	"math"
)

func main() {
	type lingkaran struct {
		x, y float64
	}
	var a, b, c lingkaran
	var luas float64
	var masuk bool
	fmt.Scan(&a.x, &a.y)
	fmt.Scan(&b.x, &b.y)
	fmt.Scan(&c.x, &c.y)
	luas = math.Sqrt(math.Pow((b.x-a.x), 2) + math.Pow((b.y-a.y), 2))
	masuk = (a.x+luas <= c.x || a.x-luas >= c.x) && (a.x+luas <= c.x || a.x-luas >= c.x)
	fmt.Println(masuk)
}
