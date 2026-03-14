package testing

func PrintIf(str string) string {
	if len(str) >= 3 {
		return "G\n"
	}
	return "Invalid input\n"
}
