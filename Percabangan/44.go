package main

import "fmt"

func main() {
	var adik, kakak int
	var menang4 bool
	fmt.Scan(&adik, &kakak)
	menang4 = kakak*2 == adik || adik == kakak
	fmt.Println(menang4)
}
