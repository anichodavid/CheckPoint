package testing

func Firstword(s string) string {
	word := ""

	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	for i < len(s) && s[i] != ' ' {
		word += string(s[i])
		i++
	}

	return word
}
