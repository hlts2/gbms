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
				ptn: "Lil",
			},
			expected: 1,
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
			str []rune
			ptn []rune
		}
		expected map[rune]int
	}{
		{
			input: struct {
				str []rune
				ptn []rune
			}{
				str: []rune("Lil and Tiny"),
				ptn: []rune("Tiny"),
			},
			expected: map[rune]int{
				'L': -1, 'i': 1, 'l': -1, ' ': -1,
				'a': -1, 'n': 2, 'd': -1,
				'T': 0, 'y': 3,
			},
		},
		{
			input: struct {
				str []rune
				ptn []rune
			}{
				str: []rune("Lil"),
				ptn: []rune("l"),
			},
			expected: map[rune]int{
				'L': -1, 'i': -1, 'l': 0,
			},
		},
		{
			input: struct {
				str []rune
				ptn []rune
			}{
				str: []rune(""),
				ptn: []rune(""),
			},
			expected: map[rune]int{},
		},
		{
			input: struct {
				str []rune
				ptn []rune
			}{
				str: []rune("a"),
				ptn: []rune("abc"),
			},
			expected: map[rune]int{
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
