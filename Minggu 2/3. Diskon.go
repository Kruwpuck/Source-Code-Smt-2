package main

import "fmt"

func main() {
	var awal int
	var member bool
	fmt.Scan(&awal, &member)
	fmt.Println(disc(awal, member))
}
func disc(awal int, member bool) int {
	if member && awal > 100000 {
		return awal * 90 / 100
	} else if awal > 100000 {
		return awal * 95 / 100
	} else {
		return awal
	}
}
