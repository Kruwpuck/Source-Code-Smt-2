package main

import "fmt"

func main() {
	var viking, saxon int
	fmt.Scan(&viking, &saxon)
	if saxon > viking*4 {
		fmt.Println("saxon menang")
	} else if saxon < viking*4 {
		fmt.Println("viking menang")
	} else {
		fmt.Println("imbang")
	}
}
