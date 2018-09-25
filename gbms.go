package gbms

import "unsafe"

// Search Searches a pattern in a text and returns count of pattern
func Search(str, ptn string) int {
	return 0
}

// toBytes convert string to bytes
func toBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}
