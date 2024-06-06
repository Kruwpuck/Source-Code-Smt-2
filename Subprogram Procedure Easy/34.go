//buatlah procedure saja
func belajarMain(aktivitas string, jumAktivitasBelajar *int, jumAktivitasMain *int){
 if aktivitas == "belajar" {
 *jumAktivitasBelajar++
 } else if aktivitas == "main" {
 *jumAktivitasMain++
 }
}