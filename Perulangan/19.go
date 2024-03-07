package main

import "fmt"

func main() {
	var turki, mongol int
	fmt.Scan(&turki, &mongol)
	if turki*3 >= mongol {
		fmt.Println("turki menang")
	} else {
		fmt.Println("mongol menang")
	}
}
