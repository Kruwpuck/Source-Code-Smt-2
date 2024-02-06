package main

import "fmt"

func main() {
	var x, y, z, pertemuan int
	fmt.Scan(&x, &y, &z)
	for i := 1; i <= 365; i++ {
		if i%x == 0 && i%y == 0 && i%z == 0 {
			pertemuan++
		}
	}
	fmt.Println(pertemuan)
}
