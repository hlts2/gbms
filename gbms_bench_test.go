package gbms

import (
	"testing"

	"github.com/cubicdaiya/bms"
)

func SearchWithBasic(str, ptn string) int {
	bStr, bPtn := toBytes(str, ptn)

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
			b.Errorf("SearchWithBasic is wrong. expected: %v, got: %v", expected, got)
		}
	}
}
