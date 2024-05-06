package main

import "fmt"

const NMax int = 1000

var episodes [NMax]Episode

type Episode struct {
	nomor  int
	durasi int
}

func main() {
	/* I.S.: Terdefinisi nilai bilangan bulat n. Data string sejumlah n buah tersedia
	   pada perangkat input. */
	/* F.S.: Array episodes diisi dengan data yang diberikan. */

	var n, i int

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&episodes[i].nomor, &episodes[i].durasi)
	}
	fmt.Println(totalDurasiEpisodeGanjil(n))
}

func totalDurasiEpisodeGanjil(m int) int {
	// Mengembalikan total durasi jika nomor episode adalah ganjil.
	var i, total int
	for i = 0; i < m; i++ {
		if episodes[i].nomor%2 != 0 {
			total += episodes[i].durasi
		}
	}
	return total
}
