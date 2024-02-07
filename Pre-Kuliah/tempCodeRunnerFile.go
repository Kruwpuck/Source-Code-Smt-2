package main

import "fmt"

func main() {
	var X, bilbuldigit4, hasiljumlah1, hasiljumlah2 int
	fmt.Scan(&X, &bilbuldigit4)
	hasiljumlah1 = (bilbuldigit4 / 1000) + (bilbuldigit4 % 10)
	hasiljumlah2 = hasiljumlah2 + ((bilbuldigit4 / 100) % 10) + ((bilbuldigit4 / 10) % 10)
	for X != hasiljumlah1 {
		fmt.Scan(&bilbuldigit4)
		hasiljumlah1 = (bilbuldigit4 / 1000) + (bilbuldigit4 % 10)
		hasiljumlah2 = hasiljumlah2 + ((bilbuldigit4 / 100) % 10) + ((bilbuldigit4 / 10) % 10)
	}
	fmt.Println(hasiljumlah2)
}
