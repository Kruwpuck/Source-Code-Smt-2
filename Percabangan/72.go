package main

import "fmt"

func main() {
	var s1, s2, s3 float64
	var hasil bool
	fmt.Scan(&s1, &s2, &s3)
	hasil = (s1+s2 > s3 && s2+s3 > s1 && s1+s3 > s2) && s3 != 0.0 && s2 != 0.0 && s1 != 0.0
	fmt.Println(hasil)
}
