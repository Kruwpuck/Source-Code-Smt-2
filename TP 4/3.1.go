package main
import "fmt"

func hitungMenang(g, k int, jm *int){
	if g>k{
		*jm+=1
	}
}
func hitungDraw(g, k int, jd *int){
	if g==k{
		*jd+=1
	}
}
func hitungKalah(g, k int, jk *int){
	if g<k{
		*jk+=1
	}
}
func hitungJumGolKegolanSelisih(g, k int, jg,jk,jsg *int){
	*jg+=g
	*jk+=k
	*jsg=*jg-*jk
}
func hitungJumPoint(jm,jd,jp *int){
	*jp=(*jm*3)+*jd
}

func main() {
	var n, gol,kegolan,jmlM,jmlD,jmlK,jmlG,jmlKG,jmlSG,jmlP,i  int
	fmt.Scan(&n)
	for i=1;i<=n;i++{
		fmt.Scan(&gol,&kegolan)
		hitungMenang(gol,kegolan,&jmlM)
		hitungDraw(gol,kegolan,&jmlD)
		hitungKalah(gol,kegolan,&jmlK)
		hitungJumGolKegolanSelisih(gol,kegolan,&jmlG,&jmlKG,&jmlSG)
	}
	hitungJumPoint(&jmlM,&jmlD,&jmlP)
	fmt.Println(n,jmlM,jmlD,jmlK,jmlG,jmlKG,jmlSG,jmlP)
}