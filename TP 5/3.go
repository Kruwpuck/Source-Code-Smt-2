package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	cetak_rerata(N, 1, 0, 0)
}

func cetak_rerata(N int, i int, jumlah int, rata float64) {
	if !(i > N) {
		jumlah += i
		rata = float64(jumlah) / float64(i)
		fmt.Println(jumlah, rata)
		cetak_rerata(N, i+1, jumlah, rata)
	}

}
