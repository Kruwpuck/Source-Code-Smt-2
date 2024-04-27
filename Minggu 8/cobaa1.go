package main

import "fmt"

const NMAX int = 10

type tim struct {
	gol                                                                           [NMAX]int
	totPertandingan, totGol, totKebobolan, totMenang, totKalah, totDraw, totPoint int
}

func main() {
	var timA, timB tim
	bacaHasil(&timA, &timB)
	fmt.Println("Statistik Tim G")
	cetakHasil(timA)
	fmt.Println()
	fmt.Println("\nStatistik Tim H")
	cetakHasil(timB)
}

func bacaHasil(tA, tB *tim) {
	/*
		IS: Tim A (tA) dan tim B (tB) terdefinisi sembarang
		Proses: Membaca skor pertandingan berupa jumlah gol tim A dan tim B.
				Mengisi field-field sesuai skor kedua tim. Field-fieldnya:
				1. total pertandingan
				2. gol setiap pertandingan
				3. total gol
				4. total kebobolan
				5. total point: point 3 untuk menang, point 1 untuk draw
				6. total menang: Menang, jika gol lebih besar dari gol lawan
				7. total draw: Draw, jika gol sama dengan gol lawan
				8. total kalah: Kalah, jika gol lebih kecil dari gol lawan
				Pembacaan skor berakhir jika kedua skor bernilai negatif.
		FS: Data tim A (tA) dan data tim B (tB) berisi nilai
	*/
	var i, x, y int
	i = 0
	fmt.Scan(&x, &y)
	for x >= 0 && y >= 0 && i < NMAX {
		tA.gol[i] = x
		tB.gol[i] = y
		tA.totGol += tA.gol[i]
		tB.totGol += tB.gol[i]
		tA.totPertandingan++
		tB.totPertandingan++
		if tA.gol[i] > tB.gol[i] {
			tA.totMenang++
			tB.totKalah++
			tA.totPoint += 3
		} else if tA.gol[i] == tB.gol[i] {
			tA.totDraw++
			tB.totDraw++
			tA.totPoint += 1
			tB.totPoint += 1
		} else {
			tB.totMenang++
			tA.totKalah++
			tB.totPoint += 3
		}
		tA.totKebobolan = tB.totGol
		tB.totKebobolan = tA.totGol
		i++
		fmt.Scan(&x, &y)
	}

}

func cetakHasil(t tim) {
	/*
		IS: Data tim (t) terdefinisi
		FS: tercetak di layar statistik pertandingan tim (t) dengan format:

		Total pertandingan: <total pertandingan>
		Gol tiap pertandingan: <gol1 gol2 ... goln>
		Total menang: <total kemenangan>
		Total draw: <total draw>
		Total kalah: <total kalah>
		Total gol: <total gol>
		Total kebobolan: <total kebobolan>
		Total point: <total point>
	*/
	fmt.Println("Total pertandingan:", t.totPertandingan)
	fmt.Print("Gol tiap pertandingan: ")
	for i := 0; i < t.totPertandingan; i++ {
		fmt.Print(t.gol[i], " ")
	}
	fmt.Println()
	fmt.Println("Total menang:", t.totMenang)
	fmt.Println("Total draw:", t.totDraw)
	fmt.Println("Total kalah:", t.totKalah)
	fmt.Println("Total gol:", t.totGol)
	fmt.Println("Total kebobolan:", t.totKebobolan)
	fmt.Print("Total point: ", t.totPoint)

}
