
const N int = 20

type data struct {
	nama string
	umur int
}
type Array [N]data

var jumlah, X int
var T Array

func CariNamaBerdasarkanUmur(T Array, jumlah int, X int) string {
   var i, idx int
   i = 0
   idx = -1
   for i < jumlah && idx == -1{
       if T[i].umur == X {
           idx = i
       } 
       i ++
   }
        if idx != -1 {
        return T[idx].nama
        } else {
            return "Tidak ditemukan"
        }

}