package main

import "fmt"

func main() {
	var jam, menit, detik, total int
	fmt.Scan(&jam, &menit, &detik)
	total += jam * 3600
	total += menit * 60
	total += detik
	fmt.Println(total)
}
