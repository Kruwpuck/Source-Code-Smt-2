package main

import "fmt"

const NMAX int = 1000

type tabChar [NMAX]byte

func main() {
	var karakter tabChar
	var nkar int
	bacaKarakter(&karakter, &nkar)
	cetakKarakter(karakter, nkar)

}

func bacaKarakter(k *tabChar, n *int) {
	var x byte
	i := 0
	fmt.Scanf("%c", &x)
	for x != '.' {
		k[i] = x
		i++
		fmt.Scanf("%c", &x)
	}
	*n = i
}

func cetakKarakter(k tabChar, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", k[i])
	}
}

func pisahKarakter(k tabChar, nkar int, kons, vok * tabChar, nkons)  {
	
}