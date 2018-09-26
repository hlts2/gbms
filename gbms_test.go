package gbms

import "testing"

func TestSearchByBasicAlgo(t *testing.T) {
	tests := []struct {
		input struct {
			str []byte
			ptn []byte
		}
		expected int
	}{
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("Lil and Tiny"),
				ptn: []byte("Tiny"),
			},
			expected: 1,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("Lil and Tiny and Tiny"),
				ptn: []byte("Ti"),
			},
			expected: 2,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("Lil"),
				ptn: []byte("Lil"),
			},
			expected: 1,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("Lil"),
				ptn: []byte("y"),
			},
			expected: 0,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte(""),
				ptn: []byte(""),
			},
			expected: 0,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("lil"),
				ptn: []byte(""),
			},
			expected: 0,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte(""),
				ptn: []byte("lil"),
			},
			expected: 0,
		},
	}

	for i, test := range tests {
		got := searchByBasicAlgo(test.input.str, test.input.ptn)

		if test.expected != got {
			t.Errorf("tests[%d] - searchByBasic is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func TestSearchByBmsAlgo(t *testing.T) {
	tests := []struct {
		input struct {
			str []byte
			ptn []byte
		}
		expected int
	}{
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("Lil and Tiny"),
				ptn: []byte("Tiny"),
			},
			expected: 1,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("Lil and Tiny and Tiny"),
				ptn: []byte("Ti"),
			},
			expected: 2,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("Lil"),
				ptn: []byte("Lil"),
			},
			expected: 1,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("Lil"),
				ptn: []byte("y"),
			},
			expected: 0,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte(""),
				ptn: []byte(""),
			},
			expected: 0,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte("lil"),
				ptn: []byte(""),
			},
			expected: 0,
		},
		{
			input: struct {
				str []byte
				ptn []byte
			}{
				str: []byte(""),
				ptn: []byte("lil"),
			},
			expected: 0,
		},
	}

	for i, test := range tests {
		got := searchByBmsAlgo(test.input.str, test.input.ptn)

		if test.expected != got {
			t.Errorf("tests[%d] - searchByBms is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}

// func TestBmsTable(t *testing.T) {
// 	tests := []struct {
// 		input struct {
// 			str []byte
// 			ptn []byte
// 		}
// 		expected map[byte]int
// 	}{
// 		{
// 			input: struct {
// 				str []byte
// 				ptn []byte
// 			}{
// 				str: []byte("Lil and Tiny"),
// 				ptn: []byte("Tiny"),
// 			},
// 			expected: map[byte]int{
// 				'L': -1, 'i': 1, 'l': -1, ' ': -1,
// 				'a': -1, 'n': 2, 'd': -1,
// 				'T': 0, 'y': 3,
// 			},
// 		},
// 		{
// 			input: struct {
// 				str []byte
// 				ptn []byte
// 			}{
// 				str: []byte("Lil"),
// 				ptn: []byte("l"),
// 			},
// 			expected: map[byte]int{
// 				'L': -1, 'i': -1, 'l': 0,
// 			},
// 		},
// 		{
// 			input: struct {
// 				str []byte
// 				ptn []byte
// 			}{
// 				str: []byte(""),
// 				ptn: []byte(""),
// 			},
// 			expected: map[byte]int{},
// 		},
// 		{
// 			input: struct {
// 				str []byte
// 				ptn []byte
// 			}{
// 				str: []byte("a"),
// 				ptn: []byte("abc"),
// 			},
// 			expected: map[byte]int{
// 				'a': 0,
// 			},
// 		},
// 	}
//
// 	for i, test := range tests {
// 		got := skipTable(test.input.str, test.input.ptn)
//
// 		if !reflect.DeepEqual(test.expected, got) {
// 			t.Errorf("tests[%d] - bmsTable is wrong. expected: %v, got: %v", i, test.expected, got)
// 		}
// 	}
// }
