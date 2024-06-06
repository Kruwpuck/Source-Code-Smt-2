package main

import "fmt"

func main() {
	var jam, p int
	fmt.Scan(&p, &jam)
	if p >= 1000 && jam >= 4000 {
		fmt.Println("sudah dapat dimonetisasi")
	} else {
		fmt.Println("belum dapat dimonetisasi")
	}
}
