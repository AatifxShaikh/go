package iterations

func Repeat(char string) string {
	result := ""
	for i := 0; i < 5; i++ {
		result += char
	}
	return result
}
