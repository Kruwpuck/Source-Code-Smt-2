type mahasiswa struct {
	nama, indeks string
	nilai        int
}

const NMAX int = 1024

type dataMahasiswa [NMAX]mahasiswa

func isiArray(himpunan *dataMahasiswa, n int) {
	// algoritma untuk menginput data mahasiswa dalam bentuk array dan index nilai
	for i := 0; i < n; i++ {
		fmt.Scan(&himpunan[i].nama, &himpunan[i].nilai)
		if himpunan[i].nilai > 80 {
			himpunan[i].indeks = "A"
		} else if himpunan[i].nilai > 70 && himpunan[i].nilai <= 80 {
			himpunan[i].indeks = "AB"
		} else if himpunan[i].nilai > 65 && himpunan[i].nilai <= 70 {
			himpunan[i].indeks = "B"
		} else if himpunan[i].nilai > 60 && himpunan[i].nilai <= 65 {
			himpunan[i].indeks = "BC"
		} else if himpunan[i].nilai > 50 && himpunan[i].nilai <= 60 {
			himpunan[i].indeks = "C"
		} else if himpunan[i].nilai > 40 && himpunan[i].nilai <= 50 {
			himpunan[i].indeks = "D"
		} else {
			himpunan[i].indeks = "E"
		}
	}
}

func insertionSort(himpunan *dataMahasiswa, n int) {
	// algoritma insertion sort secara ascending
	for i := 1; i < n; i++ {
		temp := himpunan[i]
		j := i - 1
		for j >= 0 && himpunan[j].nilai > temp.nilai {
			himpunan[j+1] = himpunan[j]
			j--
		}
		himpunan[j+1] = temp
	}
}

func showArray(himpunan dataMahasiswa, n int, tampilIndeks string) {
	// algoritma untuk menampilkan data menggunakan pengulangan array
	for i := 0; i < n; i++ {
		if himpunan[i].indeks == tampilIndeks {
			fmt.Println(himpunan[i].nama, himpunan[i].nilai, himpunan[i].indeks)
		}
	}
}