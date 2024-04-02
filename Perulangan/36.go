package main

import "fmt"

func main() {
	var n, i, jml int
	var mean float64
	fmt.Scan(&n)
	for n >= 0 {
		i += 1
		jml += n
		fmt.Scan(&n)
	}
	mean = float64(jml) / float64(i)
	fmt.Printf("%.2f\n", mean)
}
