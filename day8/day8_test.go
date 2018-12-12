package day8

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

type Node struct {
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

func Test_PartOne(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day8/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	in := string(dat)

	sum := SumAllMedataFromTree(in)
	if sum != 138 {
		t.Errorf("Got %v, expected %v", sum, 138)
	}
}

func SumAllMedataFromTree(in string) int {
	n := ReadThreeFromString(in)

	return SumAllMetadata(n)
}

func ReadThreeFromString(in string) *Node {
	tokens := strings.Split(in," ")
	values := make([]int, 0)
	for _, t := range tokens {
		i, _ := strconv.Atoi(t)
		values = append(values, i)
	}

	n, rest := CreateNodeFrom(values)
	if rest == nil || len(rest) == 0 {
		return n
	}
	return n
}

func CreateNodeFrom(in []int) (*Node, []int) {
	metadata := make([]int,0)
	childData := make([]int,0)
	rest := make([]int,0)
	if len(in) == 0 {
		return nil, nil
	}

	numberOfChildren := in[0]
	numberOfMetadata := in[1]

	if numberOfChildren == 0 {
		metadata = in[2:2+numberOfMetadata]
		if len(in) > 2+numberOfMetadata+1 {
			rest = in[2+numberOfMetadata:]
		}
	} else {
		metadata = in[len(in)-numberOfMetadata:]
	}

	childData = in[2:len(in)-numberOfMetadata]

	return &Node{
		Children: CreateFromStream(childData, numberOfChildren),
		MetadataEntries: metadata,
	}, rest
}

func CreateFromStream(in []int, childs int) []*Node {
	nodes := make([]*Node,0)
	rest := in

	for i:=0; i< childs; i++ {
		var node *Node
		node, rest =  CreateNodeFrom(rest)
		nodes = append(nodes, node)
	}

	return nodes
}

func SumAllMetadata(n *Node) int {
	localMetadataSum := 0
	for _, v := range n.MetadataEntries {
		localMetadataSum += v
	}

	if n.Children == nil || len(n.Children) == 0 {
		return localMetadataSum
	}

	sumMetadataChildren := 0
	for _, c := range n.Children {
		sumMetadataChildren += SumAllMetadata(c)
	}

	return localMetadataSum + sumMetadataChildren
}

