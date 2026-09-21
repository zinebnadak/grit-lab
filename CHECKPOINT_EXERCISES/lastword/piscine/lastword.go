package piscine

func LastWord(s string) string {
	start, end := 0, 0
	for i, char := range s {
		if char != ' ' {
			end = i + 1
			if i > 0 && s[i-1] == ' ' {
				start = i
			}
		}
	}
	return s[start:end] + "\n"
}