package gbms

// Search searches a pattern in a str and returns count of pattern
func Search(str, ptn string) int {
	if len(str) == 0 || len(ptn) == 0 {
		return 0
	}

	rStr, rPtn := []rune(str), []rune(ptn)

	matchedCnt := 0

	i := 0
	for i < len(rStr)-len(rPtn)+1 {
		j := len(rPtn) - 1
		for j >= 0 && rStr[i+j] == rPtn[j] {
			j--
		}

		if j == -1 {
			matchedCnt++
			i++
		} else if pos := matchedPos(rStr[i+j], rPtn); pos < j {
			i = i + (j - pos)
		} else {
			i++
		}
	}

	return matchedCnt
}

func matchedPos(char rune, ptn []rune) int {
	i := len(ptn) - 1
	for i >= 0 && char != ptn[i] {
		i--
	}

	return i
}

// func skipTable(str, ptn []rune) map[rune]int {
// 	table := make(map[rune]int)
// 	for i := 0; i < len(str); i++ {
// 		if _, ok := table[str[i]]; ok {
// 			continue
// 		}
//
// 		j := len(ptn) - 1
// 		for j >= 0 && str[i] != ptn[j] {
// 			j--
// 		}
//
// 		table[str[i]] = j
// 	}
// 	return table
// }
