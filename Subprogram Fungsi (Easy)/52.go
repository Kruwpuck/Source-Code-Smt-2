//buat fungsi saja
func mauNonton(kerja bool, genre bool, jam int) bool {
	if !kerja && jam > 19 && genre {
		return true
	} else {
		return false
	}
}