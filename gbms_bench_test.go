package gbms

import (
	"testing"

	"github.com/cubicdaiya/bms"
)

var (
	str = `zCpidUwdtQz659i3-T7xNakVZjN3VDj9k7tJLjkhiJPV8wXudzgUPQ2SdPBxmYAK-fCzWLJkBLfSwERen-CVEL3-dnTCmmkM_WbkCSUz_F2rRQyNhRwmXdrBS-jaDVgaWE
8sctA5hQcc-LQ7jF5NH2gjNBNsG8Z97nC7TWLyTfugrKRKUbi8kFjz2Db4zVmFZPi4USkkGkkY89AtX7yZdL3AmCsgtHfN5ujLJM
5SaA2QJNmNUeT-2ZTd_f9W7dtQz659i3-T7xNakVZjN3VDj9k7tJLCQBjGdtQz659i3-T7xNakVZjN3VDj9k7tJLXEX96LFFx6HLkTDaARuWJT4Ls6YFw8YAMN-UESyWWE
XHHFAY6Zh9dpd5Ph-L9-VyaaMMM_7ayNgZF4MZSU52y2ZwLJ4P48-gSHcNPKkiKCs4bsV7Gp8Z6erihCMdtQz659i3-T7xNakVZjN3VDj9k7tJLNb93_DZd3a4WyPLxECG
Dx7FN85rE4gzVMPYe_yKFzQzdtQz659i3-T7xNakVZjN3VDj9k7tJLyiBc_pS88xmVfX4iLteSUMXFR7zfp--_cTJyYacVPYQ3Zm-jAesKu5Rp-XLWLBjWdNB2C_zJCLfSdtQz659i3-T7xNakVZjN3VDj9k7tJL8xmVfX4iLteSUMXFR7zfp--_cTJyYacVPYQ3Zm-jAesKu5Rp-XLWLBjWdNB2C_zJ8xmVfX4iLteSUMXFR7zfp--_cTJyYacVPYQ3Zm-jAesKu5Rp-XLWLBjWdNB2C_zJ8xmVfX4iLteSUMXFR7zfp--_cTJyYacVPYQ3Zm-jAesKu5Rp-XLWLBjWdNB2C_zJ8xmVfX4iLteSUMXFR7zfp--_cTJyYacVPYQ3Zm-jAesKudtQz659i3-T7xNakVZjN3VDj9k7tJL5Rp-XLWLBjWdNB2C_zJ8xmVfX4iLteSUMXFR7zfp--_cTJyYacVPYQ3Zm-jAesKu5Rp-XLWLBjWdNB2C_zJ8xmVfX4iLteSUMXFR7zfp--_cTJyYacVPYQ3Zm-jAesKu5Rp-XLWLBjWdNB2C_zJ`
	ptn      = "dtQz659i3-T7xNakVZjN3VDj9k7tJL"
	expected = 7
)

func BenchmarkBmsAlgOfGbms(b *testing.B) {
	rStr := []rune(str)
	rPtn := []rune(ptn)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		got := searchByBmsAlgo(rStr, rPtn)
		if expected != got {
			b.Errorf("gbms.searchByBmsAlgo is wrong. expected: %v, got: %v", expected, got)
		}
	}
}

func BenchmarkBasicAlgOfGbms(b *testing.B) {
	rStr := []rune(str)
	rPtn := []rune(ptn)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		got := searchByBasicAlgo(rStr, rPtn)
		if expected != got {
			b.Errorf("gbms.searchByBasicAlgo is wrong. expected: %v, got: %v", expected, got)
		}
	}
}

func BenchmarkBms(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := bms.Search(str, ptn)
		if expected != got {
			b.Errorf("bms.SearchWithBasic is wrong. expected: %v, got: %v", expected, got)
		}
	}
}
