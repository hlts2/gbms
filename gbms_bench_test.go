package gbms

import (
	"testing"

	"github.com/cubicdaiya/bms"
)

func BenchmarkBmsAlgOfGbms(b *testing.B) {
	var (
		rStr     = []rune("lillillillilittlellllillillittle")
		rPtn     = []rune("little")
		expected = 2
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		got := searchByBmsAlgo(rStr, rPtn)
		if expected != got {
			b.Errorf("gbms.searchByBmsAlgo is wrong. expected: %v, got: %v", expected, got)
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
