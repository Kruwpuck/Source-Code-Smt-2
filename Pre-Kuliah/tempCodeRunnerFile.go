package main

import "fmt"

func main() {
	var bilbul int
	var kelipatan bool
	stop := false
	fmt.Scan(&bilbul)
	for !stop {
		bilbul = bilbul - 15
		if bilbul == 0 {
			kelipatan = true
		}
		if bilbul <= 0 {
			stop = true
		}
	}

	fmt.Println(kelipatan)

}
