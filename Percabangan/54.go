package main

import "fmt"

func main() {
	var hp, skor, level int
	var lanjut, w1, w2, w3, w4, over bool
	fmt.Scan(&hp, &skor, &level)
	w1 = hp != 0
	w2 = level == 1 && skor >= 1000
	w3 = level == 2 && skor >= 3000
	w4 = skor >= 7000
	lanjut = w1 && (w2 || w3 || w4)
	over = !lanjut
	fmt.Println(over)
}
