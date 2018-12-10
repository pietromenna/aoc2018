package day7

import (
	"io/ioutil"
	"sort"
	"strings"
	"testing"
)

type Node struct {
	Name string
	DependsOn []string
	Builds []string
}

func Test_PartOneExample(t *testing.T){
	in := "Step C must be finished before step A can begin.\nStep C must be finished before step F can begin.\n	Step A must be finished before step B can begin.\nStep A must be finished before step D can begin.\nStep B must be finished before step E can begin.\nStep D must be finished before step E can begin.\nStep F must be finished before step E can begin."

	if result := BuildOrder(in); result != "CABDFE" {
		t.Errorf("Incorrect result %v, expected: %v", result, "CABDFE")
	}

}

func Test_PartOne(t *testing.T){
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day7/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	in := string(dat)

	if result := BuildOrder(in); result != "BHRTWCYSELPUVZAOIJKGMFQDXN" {
		t.Errorf("Incorrect result %v, expected: %v", result, "CABDFE")
	}

}

func ProcessLine(in string) (string, string) {
	tokens := strings.Split(in, " ")
	return tokens[7], tokens[1]
}

func BuildOrder(in string) string {
	out := ""
	inLine := make(map[string]bool)
	built := make(map[string]bool)
	graph := make(map[string]*Node)
	queue := make([]string,0)
	lines := strings.Split(in,"\n")
	for _, l := range lines {
		n, d := ProcessLine(l)
		//Update Node
		if node, ok := graph[n]; !ok {
			graph[n] = &Node{n, []string{d},[]string{}}
		} else {
			node.DependsOn = append(graph[n].DependsOn, d)
		}
		//Update Dependency
		if node, ok := graph[d]; !ok {
			graph[d] = &Node{d, []string{},[]string{n}}
		} else {
			node.Builds = append(graph[d].Builds, n)
		}
	}

	for k, v := range graph {
		if len(v.DependsOn) == 0 {
			queue = append(queue, k)
			built[k] = true
		}
	}
	sort.Strings(queue)

	for len(queue) > 0 {
		node := graph[queue[0]]
		queue = queue[1:]
		allDependencies := true
		for _, d := range node.DependsOn {
			if _, ok := built[d]; !ok {
				allDependencies =  false
			}
		}

		if allDependencies {
			out += node.Name
			built[node.Name] = true
			for _, i := range node.Builds {
				if _, ok := inLine[i] ; !ok {
					inLine[i] = true
					queue = append(queue, i)
				}
			}
			sort.Strings(queue)
		} else {
			queue = append(queue, node.Name)
		}

	}

	return out
}