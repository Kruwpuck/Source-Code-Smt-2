package main

import "fmt"

func main() {
	var p1, p2, p3 string
	fmt.Scan(&p1, &p2, &p3)
	if p1 != p2 && p2 == "B" && p2 == p3 {
		fmt.Println("pemain 1 pemenang")
	} else if p2 != p1 && p1 == p3 {
		fmt.Println("pemain 2 pemenang")
	} else if p3 != p2 && p2 == p1 {
		fmt.Println("pemain 3 pemenang")
	} else {
		fmt.Println("imbang")
	}
}
