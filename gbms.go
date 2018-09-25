package gbms

import (
	"unsafe"
)

// Search Searches a pattern in a text and returns count of pattern
func Search(str, ptn string) int {
	return 0
}

// toBytes convert string to bytes
func toBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

func bmsTable(str, ptn []byte) map[byte]int {
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
