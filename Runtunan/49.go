package main

import "fmt"

func main() {
	var adik, kakak int
	var menang4 bool
	fmt.Scan(&adik, &kakak)
	menang4 = (adik-kakak)%2 == 0 || (kakak-adik)%2 == 0
	fmt.Println(menang4)
}
