package main

import "fmt"

func main() {
	var x, y, z, total, persen int
	fmt.Scan(&x, &y, &z)
	i := 1
	for i <= z {
		if i == 2 {
			total += x
			persen = total
		} else if i%2 == 0 || i%3 == 0 {
			total += persen + y
			persen = persen + y
		}
		i += 1
	}
	fmt.Println(total)
}