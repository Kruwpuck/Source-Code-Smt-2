package main

import "fmt"

func main() {
	var adik, kakak int
	var menang bool
	fmt.Scan(&adik, &kakak)
	menang = adik == kakak || adik-kakak == 3 || adik-kakak == 6 || kakak-adik == 3 || kakak-adik == 6
	fmt.Println(menang)
}
