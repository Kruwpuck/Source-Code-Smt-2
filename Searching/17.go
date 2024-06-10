const N int = 1000

func jumlahKataPalindrome(T [N]string, nKata int) int {
    var jumlahPalindrome int
    jumlahPalindrome = 0

    for i := 0; i < nKata; i++ {
        kata := T[i]
        jumlahHuruf := len(kata)
        stop := false

        for j := 0; j < jumlahHuruf/2; j++ {
            if kata[j] != kata[jumlahHuruf-1-j] {
                stop = true
                break
            }
        }

        if !stop {
            jumlahPalindrome++
        }
    }

    return jumlahPalindrome
}
