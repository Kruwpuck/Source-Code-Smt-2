package main

import "fmt"

func main() {
	var bil1, bil2, d1, d6, d23, d45 int
	fmt.Scan(&bil1, &bil2)
	d1 = bil1 / 100
	d23 = bil1 % 100
	d45 = bil2 / 10 * 10
	d6 = (bil2 % 10) * 100
	fmt.Println(d23+d6, d45+d1)

}
