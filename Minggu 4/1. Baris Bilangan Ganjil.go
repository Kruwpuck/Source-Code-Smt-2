package main

import "fmt"

func main() {
	var N, i int
	fmt.Scan(&N)
	i = 1
	rekur(N, i)
}
func rekur(N, i int) int {
	if !(N > i) {
		if i%2 != 0 {
			return i
		}
	} else {
		if i%2 != 0 {
			return i
		}
		rekur(N, i+1)
	}
	return i
}
