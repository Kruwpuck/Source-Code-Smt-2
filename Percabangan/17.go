package main

import "fmt"

func main() {
	var jam, bayar int
	var kend string
	fmt.Scan(&kend, &jam)
	if kend == "motor" {
		if jam > 1 {
			bayar = 2000 + ((jam - 1) * 500)
		} else {
			bayar = 2000
		}
	} else {
		if jam > 1 {
			bayar = 3000 + ((jam - 1) * 1000)
		} else {
			bayar = 3000
		}
	}
	fmt.Println(bayar)
}
