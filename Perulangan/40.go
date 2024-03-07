package main

import "fmt"

func main() {
	var i string
	fmt.Scan(&i)
	if i == "B" || i == "E" {
		fmt.Print("verteks dalam")
	} else if i == "A" {
		fmt.Print("akar")
	} else {
		fmt.Print("daun")
	}
}
