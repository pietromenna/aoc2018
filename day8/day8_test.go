package day8

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

type Node struct {
	NumChild int
	NumMeta int
	Children []*Node
	MetadataEntries []int
}

func Test_PartOneExample(t *testing.T) {
	in := "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

	sum := SumAllMedataFromTree(in)
	if sum != 138 {
		t.Errorf("Got %v, expected %v", sum, 138)
	}
}

func Test_PartTwoExample(t *testing.T) {
	in := "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

	sum := RootValueFromTree(in)
	if sum != 66 {
		t.Errorf("Got %v, expected %v", sum, 138)
	}
}

func Test_PartOne(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day8/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	in := string(dat)

	sum := SumAllMedataFromTree(in)
	if sum != 35911 {
		t.Errorf("Got %v, expected %v", sum, 138)
	}
}

func Test_PartTwo(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day8/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	in := string(dat)

	sum := RootValueFromTree(in)
	if sum != 17206 {
		t.Errorf("Got %v, expected %v", sum, 138)
	}
}

func SumAllMedataFromTree(in string) int {
	n := ReadThreeFromString(in)

	return SumAllMetadata(n)
}

func RootValueFromTree(in string) int {
	n := ReadThreeFromString(in)

	return Value(n)
}

func SumAllMetadata(n *Node) int {
	localMetadataSum := 0
	for _, v := range n.MetadataEntries {
		localMetadataSum += v
	}

	childMetadata := 0
	for _, c := range n.Children {
		childMetadata += SumAllMetadata(c)
	}
	return localMetadataSum + childMetadata
}

func ReadThreeFromString(in string) *Node {
	tokens := strings.Split(in," ")
	values := make([]int, 0)
	for _, t := range tokens {
		i, _ := strconv.Atoi(t)
		values = append(values, i)
	}

	n,_ := ReadTreeFrom(values)

	return n
}

func ReadTreeFrom(in []int) (*Node, []int) {
	n := &Node{}
	n.NumChild = in[0]
	n.NumMeta = in[1]
	rest := in[2:]
	n.Children = make([]*Node,0)
	for i := 0; i < n.NumChild; i++ {
		var child *Node
		child, rest = ReadTreeFrom(rest)
		n.Children = append(n.Children,child)
	}
	for i := 0; i < n.NumMeta; i++ {
		n.MetadataEntries = append(n.MetadataEntries, rest[0])
		rest = rest[1:]
	}

	return n, rest
}

func Value(n *Node) int{
	if len(n.Children) == 0 {
		sum := 0
		for _, v := range n.MetadataEntries {
			sum += v
		}
		return sum
	}

	sum := 0
	for _, v := range n.MetadataEntries {
		if v - 1< len(n.Children) && v-1 >= 0{
			sum += Value(n.Children[v - 1])
		}
	}
	return sum
}