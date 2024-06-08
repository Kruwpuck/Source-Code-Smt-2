func TOT(a, b, c byte, semuasalah *bool, totalPoin *int) {
    // Reset semuasalah to true initially
    *semuasalah = true
    
    // Check each member's answer
    if a == 'O' {
        *totalPoin += 1
        *semuasalah = false
        return
    }
    if b == 'O' {
        *totalPoin += 1
        *semuasalah = false
        return
    }
    if c == 'O' {
        *totalPoin += 1
        *semuasalah = false
        return
    }
}