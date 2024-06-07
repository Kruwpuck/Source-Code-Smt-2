
type barang struct {
	nama       string
	harga      int
	panjangNama int
}

const NMAX int = 1024

type dataBarang [NMAX]barang

func isiArray(himpunan *dataBarang, n int) {
	// Algoritma untuk menginput data himpunan barang ke dalam bentuk array 
	for i := 0; i < n; i++ {

		fmt.Scanln(&himpunan[i].nama, &himpunan[i].harga)

		himpunan[i].panjangNama = len(himpunan[i].nama)
	}
}

func insertionSort(himpunan *dataBarang, n int) {
	// Algoritma insertion sort secara ascending berdasarkan panjang nama barang
	for pass := 1; pass < n; pass++ {
		temp := himpunan[pass]
		i := pass - 1
		for i >= 0 && temp.panjangNama < himpunan[i].panjangNama {
			himpunan[i+1] = himpunan[i]
			i--
		}
		himpunan[i+1] = temp
	}
}

func showArray(himpunan dataBarang, n int) {
	// Algoritma untuk menampilkan data barang menggunakan pengulangan array 
	for i := 0; i < n; i++ {
		fmt.Println(himpunan[i].nama, himpunan[i].harga)
	}
}