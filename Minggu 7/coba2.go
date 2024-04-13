package main

import "fmt"

const NMAX int = 10

type tabInt struct {
	info [NMAX]byte
	data int
}

func main() {
	var data tabInt
	baca(&data)
	reverse(&data)
	cetak(data)
}

func baca(m *tabInt) {
	var x byte
	m.data = 0
	fmt.Scanf("%c", &x)
	for x != '.' && m.data < NMAX {
		m.info[m.data] = x
		m.data++
		fmt.Scanf("%c", &x)
	}
}

func cetak(m tabInt) {
	for i := 0; i < m.data; i++ {
		fmt.Printf("%c", m.info[i])
	}
}

func reverse(m *tabInt) {
	for i := 0; i < m.data/2; i++ {
		m.info[i], m.info[m.data-i-1] = m.info[m.data-i-1], m.info[i]
	}
}
