const N = 6
func searchVokal(T [N]byte) bool {
    vokal := map[byte]bool{
        'a': true, 'A': true,
        'e': true, 'E': true,
        'i': true, 'I': true,
        'o': true, 'O': true,
        'u': true, 'U': true,
    }

    for _, char := range T {
        if vokal[char] {
            return true
        }
    }
    return false
}