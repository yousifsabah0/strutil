package strutil

// IndexOf returns the index of the target of specified str.
// Returns -1 if there's no matched target.
// Returns 0 if len tar is 0.
func IndexOf(str, tar string) int {
	if len(tar) == 0 {
		return 0
	}

	for i := 0; i < len(str); i++ {
		if string(str[i]) == tar {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the last occurrence index of specified in str.
// Returns -1 if there's no matched target.
// Returns 0 if len tar is 0.
func LastIndexOf(str, tar string) int {
	if len(tar) == 0 {
		return 0
	}

	current := -1
	for i := 0; i < len(str); i++ {
		if string(str[i]) == tar {
			current = i
		}
	}
	return current
}

// StartsWith determines whether a string begins with the characters of a specified string.
// Returns true or false.
func StartsWith(str, char string) bool {
	return string(str[0]) == char
}

// EndsWith determines whether a string end with the characters of a specified string.
// Returns true or false.
func EndsWith(str, char string) bool {
	return string(str[len(str)-1]) == char
}

// Concat concatenates the string arguments to the calling string and returns a new string.
// Todo: Currently it's only supporting concatenates one word. to make it work better, I need
// to make it concatenates multiple words.
func Concat(str, words string) string {
	return str + words
}
