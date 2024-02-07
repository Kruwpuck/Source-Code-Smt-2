package main

import "fmt"

func main() {
	var adik, kakak int
	var adik_menang, menang1, menang2, menang3, menang4 bool
	fmt.Scan(&adik, &kakak)
	menang1 = adik%2 != 0 && kakak%2 == 0
	menang2 = adik%2 == 0 && kakak%2 != 0
	menang3 = kakak%adik == 0
	menang4 = adik-kakak == 0 || adik-kakak == 3 || kakak-adik == 0 || kakak-adik == 3
	adik_menang = menang1 || menang2 || menang3 || menang4
	fmt.Println(adik_menang)
}
