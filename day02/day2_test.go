package day02

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

type Ct struct {
	two   int
	three int
}

func Test_CountTwosOrThrees(t *testing.T) {
	var tests = []struct {
		in  string
		out Ct
	}{
		{
			"abcdef",
			Ct{two: 0, three: 0},
		},
		{
			"bababc",
			Ct{1, 1},
		},
		{
			"abbcde",
			Ct{1, 0},
		},
		{
			"abbcde",
			Ct{1, 0},
		},
		{
			"abcccd",
			Ct{0, 1},
		},
		{
			"aabcdd",
			Ct{2, 0},
		},
		{
			"abcdee",
			Ct{1, 0},
		},
		{
			"ababab",
			Ct{0, 2},
		},
	}

	for _, tc := range tests {
		result := Count(tc.in)
		if result.two != tc.out.two || result.three != tc.out.three {
			t.Errorf("%v is different to %v", result, tc)
		}
	}

}

func Test_Part1ChecksumExample(t *testing.T) {
	example := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}

	if CheckSum(example) != 12 {
		t.Fail()
	}
}

type CtUnique struct {
}

func Test_Part1(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day02/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	inputInString := string(dat)
	input := strings.Split(inputInString, "\n")

	fmt.Printf("%d", CheckSum(input))
	t.Fail()
}

func CheckSum(strings []string) int {
	two := 0
	three := 0

	for _, s := range strings {
		r := Count(s)
		if r.two > 0 {
			two += 1
		}
		if r.three > 0 {
			three += 1
		}
	}
	return two * three
}

func Count(s string) (result Ct) {
	result.two = 0
	result.three = 0
	counter := make(map[rune]int)
	for _, c := range s {
		if _, ok := counter[c]; !ok {
			counter[c] = 1
		} else {
			counter[c] += 1
		}
	}

	for _, v := range counter {
		if v == 2 {
			result.two += 1
		}
		if v == 3 {
			result.three += 1
		}
	}
	return
}

func CountUnique(in string) int {
	unique := 0
	counter := make(map[rune]int)
	for _, c := range in {
		if _, ok := counter[c]; !ok {
			counter[c] = 1
		} else {
			counter[c] += 1
		}
	}

	for _, v := range counter {
		if v == 1 {
			unique += 1
		}
	}
	return unique
}
func Test_Part2Example(t *testing.T) {
	example := []string{"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz"}

	for _, outer := range example {
		for _, inner := range example {
			if inner != outer {
				if CountUnique(outer+inner) == 2 {
					fmt.Printf("%s and %s \n", outer, inner)
					fmt.Println(AnsFromStrings(outer, inner))
					t.Fail()
				}
			}
		}
	}
	t.Fail()
}

func Test_Part2(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day02/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	inputInString := string(dat)
	input := strings.Split(inputInString, "\n")

	for _, outer := range input {
		for _, inner := range input {
			if inner != outer {
				if AnsFromStrings(outer, inner) != "" {
					fmt.Printf("%s and %s \n", outer, inner)
					fmt.Println(AnsFromStrings(outer, inner))
					t.Fail()
				}
			}
		}
	}
	t.Fail()
}

func AnsFromStrings(s1, s2 string) string {
	ans := ""
	diff := 0
	for i, _ := range s1{
		if diff >= 2 {
			return ""
		}
		if s1[i] == s2[i] {
			ans = ans + string(s1[i])
		} else {
			diff += 1
		}
	}
	return ans
}

