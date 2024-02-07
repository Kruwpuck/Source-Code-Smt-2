package main

import "fmt"

func main() {
	var bil, jumlah, jumlah_bil int
	var nilai, stop bool
	stop = false
	for !stop {
		fmt.Scan(&bil, &nilai)
		if bil < 0 && nilai == false {
			stop = true
		} else {
			jumlah += bil
			jumlah_bil++
		}
	}
	fmt.Println(jumlah, jumlah_bil)
}
