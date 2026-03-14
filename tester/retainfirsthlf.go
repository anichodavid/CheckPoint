package testing

func RetainFirstHalf(str string) string {
	length := len(str)

	if length == 1 {
		return str
	}
	half := length / 2

	return str[:half]
}
