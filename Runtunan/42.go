package main

import "fmt"

func main() {
	var huruf string
	var hasil1, hasil2 bool
	fmt.Scan(&huruf)
	hasil1 = huruf != "a" && huruf != "i"
	hasil2 = huruf != "u" && huruf != "e" && huruf != "o"
	hasil3 := hasil1 && hasil2
	fmt.Println(hasil3)
}
