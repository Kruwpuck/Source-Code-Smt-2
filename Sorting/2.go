type pemain struct {
	nama        string
	poin, panjangNama int
}
const NMAX int = 1024

type dataPemain [NMAX]pemain

func isiArray(himpunan *dataPemain, n int) {
	// Algoritma untuk menginput data pemain dalam bentuk array
	for i := 0; i < n; i++ {
		fmt.Scanln(&himpunan[i].nama, &himpunan[i].poin)
		himpunan[i].panjangNama = len(himpunan[i].nama)
	}
}

func selectionSort(himpunan *dataPemain, n int) {
	// Algoritma selection sort secara descending berdasarkan panjang nama pemain
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if himpunan[j].panjangNama > himpunan[maxIdx].panjangNama {
				maxIdx = j
			}
		}
		// Swap
		himpunan[i], himpunan[maxIdx] = himpunan[maxIdx], himpunan[i]
	}
}

func showArray(himpunan dataPemain, n int) {
	// Algoritma untuk menampilkan data array
	for i := 0; i < n; i++ {
		fmt.Println(himpunan[i].nama, himpunan[i].poin)
	}
}