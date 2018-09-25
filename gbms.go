package gbms

import "unsafe"

// Search Searches a needle in a haystack and returns count of needle
func Search(str, ptn string) int {
	return 0
}

// toBytes convert string to bytes
func toBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}
