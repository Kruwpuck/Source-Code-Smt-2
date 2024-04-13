package main

import "fmt"

func main() {
	var a, b bool
	fmt.Scan(&a, &b)
	if a || b {
		fmt.Print("berlibur")
	} else {
		fmt.Print("tidak berlibur")
	}
}
