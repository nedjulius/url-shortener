package controllers

import "testing"

var encodingTestCases = []struct {
	in  int
	out string
}{
	{1, "b"},
	{2, "c"},
	{1231, "3Vx"},
	{3883, "baN"},
	{1839113, "hSBh"},
}

func testEncoding(t *testing.T) {
	for _, testCase := range encodingTestCases {
		res := encode(testCase.in)
		if res != testCase.out {
			t.Errorf("Encoding is incorrect, got: %s, want %s", res, testCase.out)
		}
	}
}

var decodingTestCases = []struct {
	in  string
	out int
}{
	{"b", 1},
	{"a", 0},
	{"92", 3836},
	{"hSBh", 1839113},
}

func testDecoding(t *testing.T) {
	for _, testCase := range decodingTestCases {
		res := decode(testCase.in)
		if res != testCase.out {
			t.Errorf("Decoding is incorrect, got %d, want %d", res, testCase.out)
		}
	}
}
