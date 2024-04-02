package main

import "fmt"

func main() {
	var n, m, mtot int
	for true {
		fmt.Scan(&n, &m)
		if n != 0 && m != 0 {
			mtot += n*60 + m
		} else {
			break
		}
	}
	fmt.Print(mtot/60, mtot%60)
}
