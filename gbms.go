package gbms

import (
	"unsafe"
)

// Search searches a pattern in a text and returns count of pattern
func Search(str, ptn string) int {
	if len(str) == 0 || len(ptn) == 0 {
		return 0
	}

	bStr, bPath := toBytes(str), toBytes(ptn)
	table := skipTable(bStr, bPath)

	matchedCnt := 0

	i := 0
	for i < len(bStr)-len(bPath)+1 {
		j := len(bPath) - 1
		for j >= 0 && bStr[j+i] == bPath[j] {
			j--
		}

		if j == -1 {
			matchedCnt++
			i++
		} else if table[bStr[i+j]] < j {
			i = i + (j - table[bStr[i+j]])
		} else {
			i++
		}
	}
	return matchedCnt
}

func skipTable(str, ptn []byte) map[byte]int {
	table := make(map[byte]int, len(str))
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

// toBytes convert string to bytes
func toBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}
