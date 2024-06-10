
const NMax int = 1000

type tabSegitiga [NMax]int

func inputS(Tab1, Tab2, Tab3 *tabSegitiga, n int) {
	/*I.S. Terdefinisi bilangan bulat n.
	Proses: menerima masukan data sisi segitiga. Masukan berhenti apabila sudah
	mencapai nilai maksimal yaitu n.
	F.S. Array Tab1, Tab2, Tab3 berisi bilangan bulat yang diinputkan*/
	var i int
	for i = 1; i <= n; i++ {
		fmt.Scan(&Tab1[i], &Tab2[i], &Tab3[i])
	}
}

func CekSegitiga(Tab1, Tab2, Tab3 tabSegitiga, n int) {
	/*I.S. Terdefinisi bilangan bulat n. Kemudian Array Tab1, Tab2, dan Tab3
	yang berisi sisi segitiga sebanyak n.
	F.S. Menampilkan informasi jenis segitiga*/
	var i int
	for i = 1; i <= n; i++ {
		if Tab1[i] == Tab2[i] && Tab2[i] == Tab3[i] {
			fmt.Println("Segitiga ke", i, "adalah segitiga sama sisi")
		} else if Tab1[i] != Tab2[i] && Tab2[i] != Tab3[i] && Tab1[i] != Tab3[i] {
			fmt.Println("Segitiga ke", i, "adalah segitiga sembarang")
		} else {
			fmt.Println("Segitiga ke", i, "adalah segitiga sama kaki")
		}
	}
}
