package day3

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

var fabric [1000][1000]string

type Coord struct {
	X int
	Y int
}

type Claim struct {
	Id     int
	Left   int
	Top    int
	Width  int
	Height int
}

func ReadClaim(in string) Claim {
	out := Claim{}
	tokens := strings.Split(in," ")
	out.Id,_ = strconv.Atoi(tokens[0][1:])

	coords := strings.Split(tokens[2][:len(tokens[2])-1],",")
	out.Left, _ = strconv.Atoi(coords[0])
	out.Top, _ = strconv.Atoi(coords[1])

	size := strings.Split(tokens[3],"x")
	out.Width, _ = strconv.Atoi(size[0])
	out.Height, _ = strconv.Atoi(size[1])
	return out
}

func Test_ReadClaim(t *testing.T){
	tests := []struct{
		In string
		Claim Claim
	}{
		{
			"#1 @ 1,3: 4x4",
			Claim{1,1,3,4,4},
		},
		{
			"#2 @ 3,1: 4x4",
			Claim{2,3,1,4,4},
		},
		{
			"#3 @ 5,5: 2x2",
			Claim{3,5,5,2,2},
		},
	}

	for _, tc := range tests {
		result := ReadClaim(tc.In)
		if result != tc.Claim {
			t.Errorf("Expected: %v, Actual: %v",tc.Claim, result)
		}
	}
}

func Test_ExamplePart1(t *testing.T) {
	input := "#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"

	if zeroes := ProcessClaimsCountingDuplicates(input); zeroes != 4 {
		t.Errorf("expected: %v, actual %v", 4, zeroes)
	}
}

func Test_Part1(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day3/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	input := string(dat)
	fmt.Println(fmt.Sprintf("Zeroes: %d \n", ProcessClaimsCountingDuplicates(input)))

	t.Fail()
}

func Test_Part2(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day3/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	input := string(dat)
	fmt.Println(fmt.Sprintf("Id: %d \n", ProcessClaimsReturnUntouchedId(input)))

	t.Fail()
}

func Test_ExamplePart2(t *testing.T) {
	input := "#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"

	if untouchedId := ProcessClaimsReturnUntouchedId(input); untouchedId != 3 {
		t.Errorf("expected: %v, actual %v", 3, untouchedId)
	}
}

func ProcessClaim(c Claim) []Coord{
	coords := make([]Coord,0)
	for x := c.Left ; x < c.Left + c.Width; x++ {
		for y:= c.Top; y < c.Top + c.Height; y++ {
			coords = append(coords, Coord{x,y})
		}
	}
	return coords
}

func ProcessClaimsCountingDuplicates(in string) int {
	zeroes := 0
	fabric := make(map[Coord]int)
	lines := strings.Split(in, "\n")

	for _, line := range lines {
		//fmt.Printf("%v", line)
		c := ReadClaim(line)
		coords := ProcessClaim(c)
		for _, coord := range coords {
			if _, ok := fabric[coord]; ok {
				fabric[coord] = 0
			} else {
				fabric[coord] = c.Id
			}
		}
	}

	for _, v := range fabric {
		if v == 0 {
			zeroes += 1
		}
	}

	return zeroes
}

func ProcessClaimsReturnUntouchedId(in string) int {
	originalArea := make(map[int]int)
	afterProcessArea := make(map[int]int)
	fabric := make(map[Coord]int)
	lines := strings.Split(in, "\n")

	for _, line := range lines {
		c := ReadClaim(line)
		coords := ProcessClaim(c)
		originalArea[c.Id] = len(coords)
		for _, coord := range coords {
			if _, ok := fabric[coord]; ok {
				fabric[coord] = 0
			} else {
				fabric[coord] = c.Id
			}
		}
	}

	for _, v := range fabric {
		if v != 0 {
			if _, ok := afterProcessArea[v]; ok {
				afterProcessArea[v] += 1
			} else {
				afterProcessArea[v] = 1
			}
		}
	}

	for k, v := range afterProcessArea {
		if originalArea[k] == v {
			return k
		}
	}

	return 0
}
