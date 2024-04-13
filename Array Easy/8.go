const NMax int = 1000

type tabBuku [NMax]string

func masukanJudul(T *tabBuku, n *int) {
	/*I.S. Data string buku tersedia pada piranti masukan
	  Proses: Masukan akan berhenti ketika pengguna memasukkan string "#".
	  F.S. Kumpulan data buku sebanyak n berada dalam array T.
	  Petunjuk: gunakan while-loop dalam melakukan proses masukan
	*/
	var i int
	for i = 0; i < NMax; i++ {
		fmt.Scan(&T[i])
		if T[i] == "#" {
			break
		}
	}
	*n = i
}

func hapusJudul(T *tabBuku, n *int, idx int) {
	/*I.S. Terdefinisi n buah string judul buku pada array T dan idx yang
	  merupakan target indeks array yang akan dihapus
	  F.S. Array T berisi judul buku yang telah mengalami penghapusan pada indeks
	  tertentu, dan jumlah elemen array (n) berkurang satu.
	  Keterangan: setiap elemen yagn di sebelah kanan idx, digeser ke kiri*/
	T[idx] = ""
}

func cetakJudul(T tabBuku, n int) {
	/*I.S. Terdefinisi array T berisi n jumlah string judul buku
	  F.S. Seitap elemen pada array T ditampilkan pada layar
	*/
	for i := 0; i < n; i++ {
		if T[i] != "" {
			fmt.Print(T[i], " ")
		}
	}
	fmt.Println("")
}