package main

import "fmt"

func main() {
	var adik, kakak int
	var prima bool
	fmt.Scan(&adik, &kakak)
	prima = true
	if adik-kakak == 1 || kakak-adik == 1 {
		prima = false
	} else if adik == kakak {
		prima = false
	} else {
		if adik > kakak {
			for i := 2; i <= (adik-kakak)-1; i++ {
				if (adik-kakak)%i == 0 {
					prima = false
				}
			}
		} else {
			for i := 2; i <= (kakak-adik)-1; i++ {
				if (kakak-adik)%i == 0 {
					prima = false
				}
			}
		}
	}

	fmt.Println(prima)
}
