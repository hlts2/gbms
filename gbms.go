package gbms

// Search --
func Search(str, ptn string) int {
	// if len(str) < 100 || len(ptn) < 100 { // TODO Adjust threshold
	// 	return searchByBasicAlgo([]rune(str), []rune(ptn))
	// }
	return searchByBasicAlgo([]rune(str), []rune(ptn))
}

// SearchByBasic searches a pattern in a str by basic search algorithm and returns count of pattern
func searchByBasicAlgo(rStr, rPtn []rune) int {
	if len(rStr) == 0 || len(rPtn) == 0 {
		return 0
	}

	matchedCnt := 0

	for i := 0; i < len(rStr)-len(rPtn)+1; i++ {
		ok := true
		for j := 0; j < len(rPtn); j++ {
			if rStr[i+j] != rPtn[j] {
				ok = false
				break
			}
		}

		if ok {
			matchedCnt++
		}
	}

	return matchedCnt
}

// SearchByBms searches a pattern in a str by bms algorithm and returns count of pattern
func searchByBmsAlgo(rStr, rPtn []rune) int {
	if len(rStr) == 0 || len(rPtn) == 0 {
		return 0
	}

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
