package main

import "fmt"

func main() {
	var bil1, bil2 int
	var hasil bool

	fmt.Scan(&bil1, &bil2)
	hasil = bil1%bil2 == 0
	fmt.Println(hasil)
}
