package main

import "fmt"

const NMax int = 1000

type tabstr [NMax]string
type tabint [NMax]int

func main() {
	var n, i, totalP int
	var rata2 float64
	var nm tabstr
	var FT, TW, TH, total tabint

	// Masukkan banyaknya pemain
	fmt.Scan(&n)

	// Masukkan nama, banyaknya poin yang didapat, dan hitung total poin per pemain
	for i = 0; i < n; i++ {
		fmt.Scan(&nm[i], &FT[i], &TW[i], &TH[i])
		total[i] = FT[i]*1 + TW[i]*2 + TH[i]*3
		totalP += total[i]
	}

	// Tampilkan poin yang diperoleh per pemain dan jumlahkan seluruh total poin pemain
	for i = 0; i < n; i++ {
		fmt.Printf("Poin yang diperoleh %s adalah %d\n", nm[i], total[i])
	}

	// Hitung dan tampilkan rata-rata poin yang didapat seluruh pemain
	rata2 = float64(totalP) / float64(n)
	fmt.Printf("Rata-rata poin yang diperoleh seluruh pemain adalah %.2f\n", rata2)
}