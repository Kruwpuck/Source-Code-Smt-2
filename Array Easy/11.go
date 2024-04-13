const nMax int = 1000

type tabChar [nMax]rune

func isiArray(arr *tabChar, n *int) {
	var kata string
	fmt.Scan(&kata)
	var i int
	i = 0
	for i < len(kata) {
		arr[i] = rune(kata[i])
		i++
	}
	*n = i
}

func isSudahAda(char rune, arr tabChar, total int) bool {
	for i := 0; i < total; i++ {
		if arr[i] == char {
			return true
		}
	}
	return false
}

func isVokal(char rune) bool {
	return char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o' ||
		char == 'A' || char == 'I' || char == 'U' || char == 'E' || char == 'O'
}

func prosesHurufVokal(arr tabChar, n int, total *int, arrVokal *tabChar) {
	*total = 0
	var j int
	for i := 0; i < n; i++ {
		if isVokal(arr[i]) && !isSudahAda(arr[i], *arrVokal, *total) {
			(*arrVokal)[j] = arr[i]
			j++
			*total++
		}
	}
}

func prosesHurufKonsonan(arr tabChar, n int, total *int, arrKonsonan *tabChar) {
	*total = 0
	var j int
	for i := 0; i < n; i++ {
		if (arr[i] >= 'a' && arr[i] <= 'z' || arr[i] >= 'A' && arr[i] <= 'Z') &&
			!isVokal(arr[i]) && !isSudahAda(arr[i], *arrKonsonan, *total) {
			(*arrKonsonan)[j] = arr[i]
			j++
			*total++
		}
	}
}

func prosesHurufKarakter(arr tabChar, n int, total *int, arrKar *tabChar) {
	*total = 0
	var j int
	for i := 0; i < n; i++ {
		if !isVokal(arr[i]) && (arr[i] < 'a' || arr[i] > 'z') && (arr[i] < 'A' || arr[i] > 'Z') &&
			!isSudahAda(arr[i], *arrKar, *total) {
			(*arrKar)[j] = arr[i]
			j++
			*total++
		}
	}
}

func cetakHuruf(arr tabChar, n int) {
	var totalVokal, totalKonsonan, totalKarakterLain int
	var arrVokal, arrKonsonan, arrKarakterLain tabChar

	prosesHurufVokal(arr, n, &totalVokal, &arrVokal)
	prosesHurufKonsonan(arr, n, &totalKonsonan, &arrKonsonan)
	prosesHurufKarakter(arr, n, &totalKarakterLain, &arrKarakterLain)

	fmt.Println(totalVokal)
	if totalVokal == 0 {
		fmt.Print("Tidak ditemukan")
	} else {
		for i := 0; i < totalVokal; i++ {
			fmt.Print(string(arrVokal[i]), " ")
		}
	}
	fmt.Println("")
	fmt.Println(totalKonsonan)
	if totalKonsonan == 0 {
		fmt.Print("Tidak ditemukan")
	} else {
		for i := 0; i < totalKonsonan; i++ {
			fmt.Print(string(arrKonsonan[i]), " ")
		}
	}
	fmt.Println("")
	fmt.Println(totalKarakterLain)
	if totalKarakterLain == 0 {
		fmt.Print("Tidak ditemukan")
	} else {
		for i := 0; i < totalKarakterLain; i++ {
			fmt.Print(string(arrKarakterLain[i]), " ")
		}
	}
}