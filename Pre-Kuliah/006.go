// Program Bilangan
package main

import "fmt"

func main() {
	var kiri, kanan, jumlah int
	stop := false
	for !stop {
		fmt.Scan(&kiri, &kanan)
		if kiri%2 != 0 && kanan%2 == 0 {
			jumlah += (kiri * kanan)
		}
		if kiri == 0 && kanan == 0 {
			stop = true
		}
	}
	fmt.Println(jumlah)
}
