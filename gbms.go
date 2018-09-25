package gbms

import "unsafe"

// Search Searches a needle in a haystack and returns count of needle
func Search(str, ptn string) int {
	return 0
}

func toBytes(str1, str2 string) ([]byte, []byte) {
	return *(*[]byte)(unsafe.Pointer(&str1)), *(*[]byte)(unsafe.Pointer(&str2))
}
