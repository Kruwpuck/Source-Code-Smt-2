func cekSuara(suara_burung string, n int) bool {
	var ja, be, ci, je, lo int
	n = 20
	for i:=1;i<= n ;i++ {
	fmt.Scan(&suara_burung)
	if suara_burung == "JA" {
	ja = 2
	} else if suara_burung == "BE" {
	be = 2
	} else if suara_burung =="CI" {
	ci = 2
	} else if suara_burung == "JE" {
	je = 2
	} else if suara_burung == "LO" {
	lo = 2
	}
	}
	return ja > 0 && be > 0 && ci > 0 && je> 0 && lo>0
   }