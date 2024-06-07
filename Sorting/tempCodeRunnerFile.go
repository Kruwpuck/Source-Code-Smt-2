type mahasiwa struct {
	nama, indeks string
	nilai  int
}

const NMAX int = 1024

type dataMahasiwa [NMAX]mahasiswa

func isiArray(himpunan *dataMahasiswa, n int) {
	//algoritma untuk menginput data mahasiswa dalam bentuk array dan index nilai 
	// note :  gunakan if else untuk menampilkan nilai oindex 
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&himpunan[i].nama,&himpunan[i].nilai)
		if himpunan[i].nilai > 80 {
			himpunan[i].indeks = "A"
		} himpunan[i].nilai > 70 && himpunan[i].nilai <= 80 {
			himpunan[i].indeks = "AB"
		} himpunan[i].nilai > 65 && himpunan[i].nilai <= 70 {
			himpunan[i].indeks = "B"
		} himpunan[i].nilai > 60 && himpunan[i].nilai <= 65 {
			himpunan[i].indeks = "BC"
		} himpunan[i].nilai > 50 && himpunan[i].nilai <= 60 {
			himpunan[i].indeks = "C"
		} himpunan[i].nilai > 40 && himpunan[i].nilai <= 50 {
			himpunan[i].indeks = "D"
		} himpunan[i].nilai <= 40 {
			himpunan[i].indeks = "E"
		} 
	}
}

func insertionSort(himpunan *dataMahasiswa, n int) {
	//algoritma insertion sort secara ascending
	var pass, i int
	var temp mahasiswa
	for pass = 1; pass < n; pass++ {
		temp = himpunan[pass]
		i = pass+1
		for i>0 && temp.nilai < himpunan[i-1].nilai{
			himpunan[i] = himpunan[i-1]
			i--
		}
		himpunan = temp
	}
}

func showArray(himpunan dataMahasiwa, n int, tampilIndeks string) {
	//algoritma untuk menampilkan data menggunakan pengulangan array 
	for i := 0; i < n; i++ {
		if himpunan[i].indeks == tampilIndeks{
			fmt.Println(himpunan[i].nama, himpunan[i].nilai, himpunan[i].indeks)
		}
	}
}
