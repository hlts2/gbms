package gbms

import (
	"testing"

	"github.com/cubicdaiya/bms"
)

func SearchWithBasic(str, ptn string) int {
	rStr, rPath := []rune(str), []rune(ptn)

	matchedCnt := 0

	for i := 0; i < len(rStr)-len(rPath)+1; i++ {
		ok := true
		for j := 0; j < len(rPath); j++ {
			if rStr[i+j] != rPath[j] {
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

func BenchmarkBasicPattern(b *testing.B) {
	var (
		str      = "lillillillilittlellllillillittle"
		ptn      = "little"
		expected = 2
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		got := SearchWithBasic(str, ptn)
		if expected != got {
			b.Errorf("SearchWithBasic is wrong. expected: %v, got: %v", expected, got)
		}
	}
}

func BenchmarkGbms(b *testing.B) {
	var (
		str      = "lillillillilittlellllillillittle"
		ptn      = "little"
		expected = 2
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		got := Search(str, ptn)
		if expected != got {
			b.Errorf("gbms.Search is wrong. expected: %v, got: %v", expected, got)
		}
	}
}

func BenchmarkBms(b *testing.B) {
	var (
		str      = "lillillillilittlellllillillittle"
		ptn      = "little"
		expected = 2
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		got := bms.Search(str, ptn)
		if expected != got {
			b.Errorf("bms.SearchWithBasic is wrong. expected: %v, got: %v", expected, got)
		}
	}
}

func BenchmarkSkip(b *testing.B) {
	var (
		str = []rune("lillillillilittlellllillillittle")
		ptn = []rune("little")
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = skipTable(str, ptn)
	}
}
