package main

import "fmt"

func main() {
	var quantum, galactum, nebula, sterallis, lumin int
	fmt.Scan(&lumin)
	quantum = lumin / 1200
	galactum = (lumin % 1200) / 120
	nebula = (lumin % 1200 % 120) / 40
	sterallis = (lumin % 1200 % 120 % 40) / 20
	lumin = (lumin % 1200 % 120 % 40 % 20)
	fmt.Println(quantum, galactum, nebula, sterallis, lumin)
}
