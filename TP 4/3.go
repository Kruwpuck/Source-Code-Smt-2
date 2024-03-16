package main

import "fmt"

func main() {
	var n, g, k, jm, jd, jk, jg, jb, jsg, jp int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&g, &k)
		hitungMenang(g, k, &jm)
		hitungDraw(g, k, &jd)
		hitungKalah(g, k, &jk)
		hitungJumGolKegolanSelisih(g, k, &jg, &jb, &jsg)
		hitungJumPoint(&jp, &jm, &jd)
	}
	fmt.Println(n, jm, jd, jk, jg, jb, jsg, jp)
}
func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm++
	}
}
func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}
func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk++
	}
}
func hitungJumGolKegolanSelisih(g, k int, jg, jb, jsg *int) {
	*jg += g
	*jb += k
	*jsg = *jg - *jb

}
func hitungJumPoint(jp, jm, jd *int) {
	*jp = (*jm * 3) + *jd
}
