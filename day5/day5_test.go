package day5

import (
	"io/ioutil"
	"testing"
	"unicode"
)

func Test_PartOne(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day5/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	input := string(dat)
	out := PartOne(input)
	if len([]rune(out)) != 9172 {
		t.Errorf("Got %v from %s", len([]rune(out)), out)
	}
}

func Test_PartOneExamples(t *testing.T) {
	testCases := []struct {
		in       string
		expected string
	}{
		{
			"aA",
			"",
		},
		{
			"abBA",
			"",
		},
		{
			"abAB",
			"abAB",
		},
		{
			"aabAAB",
			"aabAAB",
		},
		{
			"dabAcCaCBAcCcaDA",
			"dabCBAcaDA",
		},
	}

	for _, tc := range testCases {
		if e := PartOne(tc.in); e != tc.expected {
			t.Errorf("Expected %v, got %v from %v", tc.expected, e, tc.in)
		}
	}
}

func PartOne(i string) string {
	out, changed := React(i)
	if !changed {
		return out
	} else {
		for changed {
			out, changed = React(out)
		}
	}
	return out
}

func React(in string) (o string, changed bool) {
	out := ""
	changed = false
	inLen := len([]rune(in))
	for i := 0; i < inLen; i++ {
		if (unicode.IsUpper(rune(in[i])) && i < inLen - 1 && unicode.IsLower(rune(in[i+1])) && unicode.ToUpper(rune(in[i+1])) == rune(in[i])) ||
			(unicode.IsLower(rune(in[i])) && i < inLen - 1 && unicode.IsUpper(rune(in[i+1])) && unicode.ToLower(rune(in[i+1])) == rune(in[i])) {
			i += 1
			changed = true
		} else {
			out += string(rune(in[i]))
		}
	}
	return out, changed
}

