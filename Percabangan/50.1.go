package main

import "fmt"

func main() {
	var adik, kakak int
	var menang4 bool
	fmt.Scan(&adik, &kakak)
	menang4 = (adik-kakak != 1 && kakak-adik != 1) && adik-kakak == 2 || kakak-adik == 2 || adik-kakak == 3 || kakak-adik == 3 || adik-kakak == 5 || kakak-adik == 5 || adik-kakak == 7 || kakak-adik == 7
	fmt.Println(menang4)
}
