package main

import "fmt"

func main() {
	var bilangan int
	var hasil bool

	fmt.Scan(&bilangan)
	hasil = bilangan < 0

	fmt.Print(hasil)
}
