package testing

func Hashcode(dec string) string {
	n := len(dec)
	result := ""

	for i := 0; i < n; i++ {
		hash := (int(dec[i]) + n) % 127

		if hash < 33 {
			hash += 33
		}

		result += string(rune(hash))
	}
	return result
}
