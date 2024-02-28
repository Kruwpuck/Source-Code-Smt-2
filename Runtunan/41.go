package main

import "fmt"

func main() {
	var huruf string
	var hasil bool
	fmt.Scan(&huruf)
	hasil = huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o"
	fmt.Println(hasil)
}
