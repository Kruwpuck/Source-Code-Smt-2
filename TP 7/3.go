package main

import "fmt"

// Tipe bentukan pegawai dengan komponen (field) nama, gaji_pokok, masa_kerja, dan besar_bonus
type pegawai struct {
	nama                    string
	gaji_pokok, besar_bonus int
	masa_kerja              float64
}

func main() {
	// deklarasi variabel bertipe pegawai
	var karyawan pegawai

	// baca data pengawai
	fmt.Scan(&karyawan.nama, &karyawan.gaji_pokok, &karyawan.masa_kerja)

	// hitung bonus dengan memanggil prosedur hitung_bonus
	hitung_bonus(&karyawan)

	// Cetak besar bonus
	fmt.Println("Pegawai dengan nama", karyawan.nama, "mendapatkan bonus sebesar Rp", karyawan.besar_bonus)
}

func hitung_bonus(p *pegawai) {
	/* IS: p.nama, p.gaji_pokok, p.masa_kerja terdefinisi
	   Proses: Besar bonus dihitung dengan mengalikan masa kerja dengan angka pengali
	           Jika masa kerja minimal 10 tahun, angka pengalinya 1.5
	           jika masa kerja di bawah 10 tahun hingga 5 tahun, angka pengalinya 0.75
			   di bawah 5 tahun, angka pengalinya 0.5
	   FS: p.besar_bonus berisi nilai
	*/
	if p.masa_kerja >= 10.0 {
		p.besar_bonus = p.gaji_pokok * 150 / 100
	} else if p.masa_kerja < 10.0 && p.masa_kerja >= 5.0 {
		p.besar_bonus = p.gaji_pokok * 75 / 100
	} else {
		p.besar_bonus = p.gaji_pokok * 50 / 100
	}

}
