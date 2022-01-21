package strutil

// BasicSearch returns the pattern if matched in a str.
// Returns -1 if there's no match.
// Implemented using the Naive searching algorithm (https://en.wikipedia.org/wiki/String-searching_algorithm#Naive_string_search).
// This algorithm designed for smaller texts.
func BasicSearch(str, pattern string) int {
	var strLength int = len(str)
	var patternLength int = len(pattern)

	for i := 0; i <= strLength-patternLength; i++ {
		var j int
		for j = 0; j < patternLength; j++ {
			if str[i+j] != pattern[j] {
				break
			}
		}

		if j == patternLength {
			return i
		}
	}
	return -1
}

//TODO: Complete this function.
func Search(text []string, pattern []string) {

}
