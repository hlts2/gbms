package gbms

// Search --
func Search(str, ptn string) int {
	// if len(str) < 100 || len(ptn) < 100 { // TODO Adjust threshold
	// 	return searchByBasicAlgo([]byte(str), []byte(ptn))
	// }
	return searchByBasicAlgo([]byte(str), []byte(ptn))
}

// SearchByBasic searches a pattern in a str by basic search algorithm and returns count of pattern
func searchByBasicAlgo(bStr, bPtn []byte) int {
	if len(bStr) == 0 || len(bPtn) == 0 {
		return 0
	}

	matchedCnt := 0

	for i := 0; i < len(bStr)-len(bPtn)+1; i++ {
		ok := true
		for j := 0; j < len(bPtn); j++ {
			if bStr[i+j] != bPtn[j] {
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
func searchByBmsAlgo(bStr, bPtn []byte) int {
	if len(bStr) == 0 || len(bPtn) == 0 {
		return 0
	}

	matchedCnt := 0

	i := 0
	for i < len(bStr)-len(bPtn)+1 {
		j := len(bPtn) - 1
		for j >= 0 && bStr[i+j] == bPtn[j] {
			j--
		}

		if j == -1 {
			matchedCnt++
			i = i + len(bPtn)
		} else if pos := matchedPos(bStr[i+j], bPtn); pos < j {
			i = i + (j - pos)
		} else {
			i++
		}
	}

	return matchedCnt
}

func matchedPos(char byte, ptn []byte) int {
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
