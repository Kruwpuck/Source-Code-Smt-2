package main

import "fmt"

func main() {
	var x, y, z, boom int
	fmt.Scan(&x, &y, &z)
	i := x
	for i <= z {
		boom++
		if i+y <= z {
			i += y
			boom++
		} else {
			break
		}
		if i+x <= z {
			i += x
		} else {
			break
		}
	}
	fmt.Print(boom)
}
