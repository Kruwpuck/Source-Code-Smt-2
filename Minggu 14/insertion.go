package main

import "fmt"

// func main() {
// 	// Insetion Ascending
// 	A := [7]int{6, 8, 4, 2, 6, 9, 1}
// 	var pass, i, temp int
// 	for pass=1 ; pass <= 7-1;pass++{
// 		temp = A[pass]
// 		i = pass
// 		for i > 0 && A[i-1] > temp {
// 			A[i] = A[i-1]
// 			i--
// 		}
// 		A[i] = temp
// 	}
// 	for i := 0; i < 7; i++ {
// 		fmt.Print(A[i], " ")
// 	}
// }

// func main()  {
// 	// Selection Descending
// 	A := [7]int{6, 8, 4, 2, 7, 9, 1}
// 	var i, idx int
// 	for pass := 0; pass < 7; pass++ {
// 		idx = pass
// 		for i = pass + 1; i < 7; i++ {
// 			if A[i] > A[idx]{
// 				idx = i
// 			}
// 		}
// 		A[pass],A[idx] = A[idx],A[pass]
// 	}
// 	for i = 0; i < 7; i++ {
// 		fmt.Print(A[i], " ")
// 	}
// }

func main() {
	// Binary Descending
	A := [7]int{7,6,5,4,3,2,1}
	var left,mid,right int
	var idx int = -1
	var x int = 6
	left = 0
	right = 7 - 1
	mid = (left+right)/2
	for left <= right && idx == -1{
		if A[mid] == x {
			idx = mid
		} else if A[mid] < x {
			right = mid - 1
		}else{
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	if idx != -1{
		fmt.Println(idx)
		fmt.Println(A[idx])
	}
}