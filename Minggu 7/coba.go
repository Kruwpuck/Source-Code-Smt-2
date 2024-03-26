package main

import "fmt"

func main() {
	var i, x int
	i = 0
	x = 6
	for x < 100 {
		i++
		x++
	}
	fmt.Println(i)
}
