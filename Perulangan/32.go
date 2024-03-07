package main

import "fmt"

func main() {
	var m, c float64
	fmt.Scan(&m, &c)
	if m*c == 0 {
		fmt.Println("melewati")
	} else {
		fmt.Print("tidak melewati")
	}
}
