package strutil

import (
  "strings"
)

// findLongestWord returns the longest word in a sequence.
func FindLongestWord (str string) string {
    best, length := "", 0
    for _, s := range strings.Split(str, " ") {
        if len(s) > length {
            best, length = s, len(s)
        }
    }
    return best
}
