package main

import "fmt"

func main() {
	var bilangan int
	var hasil bool
	fmt.Scan(&bilangan)
	hasil = bilangan%2 != 0
	fmt.Println(hasil)
}
