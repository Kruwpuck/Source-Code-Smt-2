package main

import (
	"fmt"
)

func main() {
	var totalPoin int
	var pertanyaanTerjawab int

	for {
		var jawaban string
		fmt.Scanln(&jawaban)

		if len(jawaban) != 3 {
			continue
		}

		// Menghentikan jika semua jawaban adalah 'T'
		if jawaban == "TTT" {
			break
		}

		// Menghitung poin
		for _, j := range jawaban {
			if j == 'O' {
				totalPoin++
				break
			}
		}

		pertanyaanTerjawab++
		if pertanyaanTerjawab == 10 {
			break
		}
	}

	fmt.Printf("Poin = %d\n", totalPoin)
}
