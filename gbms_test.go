package gbms

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		input struct {
			str string
			ptn string
		}
		expected int
	}{
		{
			input: struct {
				str string
				ptn string
			}{
				str: "Lil and Tiny",
				ptn: "Tiny",
			},
			expected: 1,
		},
		{
			input: struct {
				str string
				ptn string
			}{
				str: "Lil and Tiny and Tiny",
				ptn: "Ti",
			},
			expected: 2,
		},
		{
			input: struct {
				str string
				ptn string
			}{
				str: "Lil",
				ptn: "y",
			},
			expected: 0,
		},
		{
			input: struct {
				str string
				ptn string
			}{
				str: "",
				ptn: "",
			},
			expected: 0,
		},
		{
			input: struct {
				str string
				ptn string
			}{
				str: "",
				ptn: "ti",
			},
			expected: 0,
		},
	}

	for i, test := range tests {
		got := Search(test.input.str, test.input.ptn)

		if test.expected != got {
			t.Errorf("tests[%d] - Search is wrong. expected: %v, got: %v", i, test.expected, got)
		}
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
		got := skipTable(test.input.str, test.input.ptn)

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("tests[%d] - bmsTable is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func TestToBytes(t *testing.T) {
	tests := []struct {
		input struct {
			str string
		}
		expected []byte
	}{
		{
			input: struct {
				str string
			}{
				str: "lil",
			},
			expected: []byte("lil"),
		},
		{
			input: struct {
				str string
			}{
				str: "",
			},
			expected: []byte(""),
		},
	}

	for i, test := range tests {
		got := toBytes(test.input.str)

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("tests[%d] - toBytes is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}
