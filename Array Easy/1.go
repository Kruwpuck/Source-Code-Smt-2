package main

import "fmt"

const NMax = 1000

func jumlahBilangan(bilangan [NMax]int) int {
	var jumlah int
	jumlah = 0
	for i := 0; i < NMax; i++ {
		fmt.Scan(&bilangan[i])
		if bilangan[i] == 0 {
			break
		}
		jumlah += bilangan[i]
	}
	return jumlah
}

func main() {
	var bilangan [NMax]int
	fmt.Println(jumlahBilangan(bilangan))
}
