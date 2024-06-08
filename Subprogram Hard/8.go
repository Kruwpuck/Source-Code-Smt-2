package main

import "fmt"

// Fungsi kalkulator
func kalkulator(a, b float64, operator byte) (float64, error) {
	switch operator {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b != 0 {
			return a / b, nil
		} else {
			return 0, fmt.Errorf("Error: Pembagian dengan nol")
		}
	default:
		return 0, fmt.Errorf("Error: Operator tidak valid")
	}
}

func main() {
	var a, b float64
	var operator byte

	// Membaca input dari pengguna
	fmt.Scanf("%f %f %c", &a, &b, &operator)

	// Memanggil fungsi kalkulator dan menyimpan hasilnya
	hasil, err := kalkulator(a, b, operator)

	// Memeriksa apakah ada error
	if err != nil {
		fmt.Println(err)
	} else {
		// Menampilkan hasil dengan pembulatan 2 digit di belakang koma
		fmt.Printf("%.2f\n", hasil)
	}
}
