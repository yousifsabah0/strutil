package strutil

// CharAt returns the character at the specified position.
// Returns an empty string if the length of str == 0 or if the position is negative.
// Returns the str if the str is only a one character.
func CharAt(str string, pos int) string {
	switch {
	case len(str) == 0:
		return ""
	case len(str) == 1:
		return str
	case pos <= 0:
		return ""
	default:
		return string(str[pos])
	}
}

// CharCodeAt returns a number indicating the Unicode value of the character at the specified position.
// Returns -1 if length of str is 0.
func CharCodeAt(str string, pos int) rune {
	if len(str) == 0 {
		return -1
	}

	r := []rune(str)
	return r[pos]
}
