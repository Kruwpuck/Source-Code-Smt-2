package main

import "fmt"

func main() {
	var adik, kakak int
	var adik_menang, menang4 bool
	fmt.Scan(&adik, &kakak)
	menang4 = adik-kakak == 1 || adik-kakak == 5 || kakak-adik == 1 || kakak-adik == 5 || adik == kakak
	adik_menang = menang4
	fmt.Println(adik_menang)
}
