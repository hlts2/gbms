package gbms

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func TestBmsTable(t *testing.T) {
	tests := []struct {
		input struct {
			str []byte
			ptn []byte
		}
		expected map[byte]int
	}{
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("Lil and Tiny"),
				ptn: []byte("Tiny"),
			},
			expected: map[byte]int{
				'L': -1, 'i': 1, 'l': -1, ' ': -1,
				'a': -1, 'n': 2, 'd': -1,
				'T': 0, 'y': 3,
			},
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("Lil"),
				ptn: []byte("l"),
			},
			expected: map[byte]int{
				'L': -1, 'i': -1, 'l': 0,
			},
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte(""),
				ptn: []byte(""),
			},
			expected: map[byte]int{},
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("a"),
				ptn: []byte("abc"),
			},
			expected: map[byte]int{
				'a': 0,
			},
		},
	}

	for i, test := range tests {
		got := bmsTable(test.input.str, test.input.ptn)

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("tests[%d] - bmsTable is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}
