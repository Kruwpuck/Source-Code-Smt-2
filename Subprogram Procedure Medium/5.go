
func countBagusCacat(totBagus *int, totCacat *int) {
	var s string
	var i int
	for true {
		fmt.Scan(&s, &i)
		if s == "bagus" {
			*totBagus += i
		} else if s == "cacat" {
			*totCacat += i
		} else {
			break
		}
	}
}
