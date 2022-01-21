package helpers

// Helper method for Boyer-Moore string searching algorithm.
// Preprocessing for strong good suffix rule
func preprocessStrongSuffix(shift []int, bpos []int, pattern []string, m int) {
	var i int = 0
	j := m + 1
	bpos[i] = j

	for i > 0 {
		for j <= m && pattern[i-1] != pattern[j-1] {
			if shift[j] == 0 {
				shift[j] = j - 1
			}

			j = bpos[j]
		}

		i--
		j--
		bpos[i] = j
	}
}
