const nMax int = 1000
type tabInt [nMax]int
func isiArray(arr1, arr2 *tabInt, n1, n2 int) {
	/* I.S. Data tersedia dalam piranti masukan
	F.S. arr1, arr2 telah terisi sejumlah bilangan */
	var bil int
	i := 0
	for i < n1 {
		fmt.Scan(&bil)
		if bil > 0 {
			arr1[i] = bil
			i++
		}
	}

	i = 0
	for i < n2 {
		fmt.Scan(&bil)
		if bil > 0 {
			arr2[i] = bil
			i++
		}
	}
}

func cekBilangan(arr tabInt, n, x int) bool {
	/*mengembalikan true apabila x ditemukan didalam array array arr*/
	for i := 0; i < n; i++ {
		if arr[i] == x {
			return true
		}
	}
	return false
}

func cetakIntersection(arr1, arr2 tabInt, n1, n2 int) {
	/* I.S. array arr1,arr2 berisi sejumlah n bilangan bulat positif
	F.S. bilangan-bilangan yang merupakan intersection dari elemen array 1 dan array 2 ditampilkan di layar */

	var found bool
	for i := 0; i < n1; i++ {
		if cekBilangan(arr2, n2, arr1[i]) {
			fmt.Print(arr1[i], " ")
			found = true
		}
	}

	if !found {
		fmt.Print("-")
	}
}