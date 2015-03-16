package slice

func All(n int, s string) (answers []string) {

	for i := 0; i+n <= len(s); i++ {
		answers = append(answers, s[i:i+n])
	}

	return answers
}

func First(n int, s string) (slice string, ok bool) {
	if n <= len(s) {
		return s[:n], true
	}

	return "", false
}

func Frist(n int, s string) string {
	frist, _ := First(n, s)
	return frist
}
