package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	if a == "bandung" && (b == "tasikmalaya" || b == "cianjur") {
		fmt.Println("bertetanggaan")
	} else if a == "tasikmalaya" && b == "bandung" {
		fmt.Println("bertetanggaan")
	} else if a == "cianjur" && (b == "bogor" || b == "bandung" || b == "sukabumi") {
		fmt.Println("bertetanggaan")
	} else if a == "sukabumi" && (b == "bogor" || b == "cianjur") {
		fmt.Println("bertetanggaan")
	} else if a == "bogor" && (b == "sukabumi" || b == "cianjur") {
		fmt.Println("bertetanggaan")
	} else {
		fmt.Println("tidak bertetanggaan")
	}
}
