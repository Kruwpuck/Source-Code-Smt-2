package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanln(&x, &y)

	// Menghitung jumlah hari pertemuan yang merupakan kelipatan x dan kelipatan y
	totalMeetings := 0
	for day := 1; day <= 365; day++ {
		if day%x == 0 && day%y == 0 {
			totalMeetings++
		}
	}

	fmt.Println(totalMeetings)
}