func searchKonsonan(x string) bool {
    konsonan := map[rune]bool{
        'b': true, 'B': true,
        'c': true, 'C': true,
        'd': true, 'D': true,
        'f': true, 'F': true,
        'g': true, 'G': true,
        'h': true, 'H': true,
        'j': true, 'J': true,
        'k': true, 'K': true,
        'l': true, 'L': true,
        'm': true, 'M': true,
        'n': true, 'N': true,
        'p': true, 'P': true,
        'q': true, 'Q': true,
        'r': true, 'R': true,
        's': true, 'S': true,
        't': true, 'T': true,
        'v': true, 'V': true,
        'w': true, 'W': true,
        'x': true, 'X': true,
        'y': true, 'Y': true,
        'z': true, 'Z': true,
    }

    for _, char := range x {
        if konsonan[char] {
            return true
        }
    }
    return false
}
