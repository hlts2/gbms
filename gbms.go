package gbms

// Search searches a pattern in a text and returns count of pattern
func Search(str, ptn string) int {
	if len(str) == 0 || len(ptn) == 0 {
		return 0
	}

	rStr, rPath := []rune(str), []rune(ptn)
	table := skipTable(rStr, rPath)

	matchedCnt := 0

	i := 0
	for i < len(rStr)-len(rPath)+1 {
		j := len(rPath) - 1
		for j >= 0 && rStr[i+j] == rPath[j] {
			j--
		}

		if j == -1 {
			matchedCnt++
			i++
		} else if pos := table[rStr[i+j]]; pos < j {
			i = i + (j - pos)
		} else {
			i++
		}
	}

	return matchedCnt
}

func skipTable(str, ptn []rune) map[rune]int {
	table := make(map[rune]int, len(str))
	for i := 0; i < len(str); i++ {
		j := len(ptn) - 1
		for j >= 0 {
			if str[i] == ptn[j] {
				break
			}
			j--
		}
		table[str[i]] = j
	}
	return table
}
