package piscine

func LastWord(s string) string {
	word := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' { // has to be a letter
			if i > 0 && s[i-1] == ' ' {  // starting of a new word eg. previous char was a space. Beginning of the string (0) does not have a previous char worth checking
				word = "" // clear when starting of a new word
			}
			word += string(s[i]) // add letter
		}
	}
	return word + "\n"
}