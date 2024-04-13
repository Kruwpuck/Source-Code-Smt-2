package main

import "fmt"

func main() {
	var masuk string
	var a, b, c, d int
	for true {
		fmt.Scan(&masuk, &a)
		if masuk == "x" && a == 0 {
			break
		} else if masuk == "sd" {
			b += a
		} else if masuk == "smp" {
			c += a
		} else {
			d += a
		}
	}
	fmt.Print(b, c, d)
}
