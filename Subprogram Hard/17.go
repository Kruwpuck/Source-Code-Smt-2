func skorTertinggiPanahan(s1, s2, s3 int, t1, t2, t3, maks *int) {
	// Tambahkan skor masing-masing pemanah ke total skor mereka
	*t1 += s1
	*t2 += s2
	*t3 += s3

	// Tentukan skor tertinggi dari ketiga total skor
	*maks = *t1
	if *t2 > *maks {
		*maks = *t2
	}
	if *t3 > *maks {
		*maks = *t3
	}
}