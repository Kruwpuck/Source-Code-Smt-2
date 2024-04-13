const NMax int = 1000

type tabString [NMax]string

func masukanArray(T *tabString, n int) {
	/*I.S. Terdefinisi nilai bilangan bulat n. Data string sebanyak n buah tersedia
	  pada input device*/
	/*F.S. Array T berisi data yang diberikan*/
	/*Petunjuk : Lakukan perulangan sebanyak n kali untuk melakukan input terhadap
	  data nilai mahasiswa*/
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}

func isUniform(T tabString, n int) bool {
	/*Mengembalikan boolean true jika isi data pada array T adalah seragam, dan
	  false jika tidak*/
	if n == 1 {
		return true
	} else {
		hasil := false
		for i := 1; i < n; i++ {
			hasil = T[i-1] == T[i]
		}
		return hasil
	}
}
