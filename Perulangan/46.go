package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a < 0 {
		fmt.Print("Freezing weather")
	} else if a >= 0 && a <= 9 {
		fmt.Print("Very Cold weather")
	} else if a >= 10 && a <= 19 {
		fmt.Print("Cold weather")
	} else if a >= 20 && a <= 29 {
		fmt.Print("Normal in Temp")
	} else if a >= 30 && a <= 39 {
		fmt.Print("It's Hot")
	} else if a >= 40 {
		fmt.Print("It's Very Hot")
	}
}
