package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var moves string
	fmt.Scan(&moves)

	var winsA, winsB int

	for i := 0; i < n; i++ {
		moveA := moves[2*i]
		moveB := moves[2*i+1]

		if (moveA == 'g' && moveB == 'k') || (moveA == 'k' && moveB == 'b') || (moveA == 'b' && moveB == 'g') {
			winsA++
		} else if (moveB == 'g' && moveA == 'k') || (moveB == 'k' && moveA == 'b') || (moveB == 'b' && moveA == 'g') {
			winsB++
		}
	}

	fmt.Printf("%d %d\n", winsA, winsB)
}