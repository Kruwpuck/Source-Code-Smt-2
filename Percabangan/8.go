package main

import "fmt"

func main() {
	var w1, w2 string
	fmt.Scan(&w1, &w2)
	if (w1 == "merah" && w2 == "kuning") || (w1 == "kuning" && w2 == "merah") {
		fmt.Println("orange")
	} else if (w1 == "biru" && w2 == "kuning") || (w1 == "kuning" && w2 == "biru") {
		fmt.Println("hijau")
	} else {
		fmt.Println("ungu")
	}
}
