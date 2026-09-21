package piscine

func LastWord(s string) string {
	end := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' { // s[i] is the character, i is its index
			if end == 0 { // still no word found yet
				end = i + 1 // ending index
			}
		} else if end != 0 { // a space after the word: its start is i+1
			return s[i+1:end] + "\n"
		}
	}
	return s[:end] + "\n" // no space found after the word, so it starts at 0
}