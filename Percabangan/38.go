package main

import "fmt"

func main() {
	var total, x float64
	for i := 0; i < 6; i++ {
		fmt.Scan(&x)
		total += x
	}
	if total/6 >= 3.33 {
		fmt.Print("mendapat sertifikat")
	} else {
		fmt.Print("tidak mendapat sertifikat")
	}
}
