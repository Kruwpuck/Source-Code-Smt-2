package main

import "fmt"

func main() {
	var x string
	for true {
		fmt.Scan(&x)
		if x == "belum" {
			fmt.Println("cari beasiswa")
		} else {
			fmt.Print("pencarian selesai")
			break
		}
	}
}
