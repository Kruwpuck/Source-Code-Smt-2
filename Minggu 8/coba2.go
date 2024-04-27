package main

import "fmt"

const NMAX int = 10

type tabStr [NMAX]string

func main() {
	data := tabStr{"sepatu", "payung", "kemeja", "jam"}
	var n int
	n = 4
	var x string
	fmt.Scan(&x)
	fmt.Println(seq_search(data, n, x))
	fmt.Println(seq_search_idx(data, n, x))
}
func seq_search(A tabStr, N int, X string) bool {
	var i int
	var ketemu bool
	i = 0
	ketemu = false
	for i < N && !ketemu {
		ketemu = A[i] == X
		i++
	}
	return ketemu
}
func seq_search_idx(A tabStr, N int, X string) int {
	var i, idx int
	idx = -1
	i = 0
	for i < N && idx == -1 {
		idx = i
		i++
	}
	return idx
}
