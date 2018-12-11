package day7

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"testing"
)

type Node struct {
	Name      string
	DependsOn []string
	Builds    []string
}

type Worker struct {
	Duration          int
	Node              string
	ExtraStepDuration int
}

func (w *Worker) IsFree() bool {
	return w.Duration <= 0
}

func (w *Worker) Tick() {
	if w.Duration > 0 {
		w.Duration = w.Duration - 1
	}
}

func (w *Worker) AssignWork(nodeName string) {
	w.Node = nodeName
	w.Duration = int(nodeName[0]-'A'+1) + w.ExtraStepDuration
}

func Test_PartOneExample(t *testing.T) {
	in := "Step C must be finished before step A can begin.\nStep C must be finished before step F can begin.\n	Step A must be finished before step B can begin.\nStep A must be finished before step D can begin.\nStep B must be finished before step E can begin.\nStep D must be finished before step E can begin.\nStep F must be finished before step E can begin."

	if result := BuildOrder(in); result != "CABDFE" {
		t.Errorf("Incorrect result %v, expected: %v", result, "CABDFE")
	}

}

func Test_PartOne(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day7/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	in := string(dat)

	if result := BuildOrder(in); result != "BHRTWCYSELPUVZAOIJKGMFQDXN" {
		t.Errorf("Incorrect result %v, expected: %v", result, "CABDFE")
	}
}

func Test_PartTwoExample(t *testing.T) {
	in := "Step C must be finished before step A can begin.\nStep C must be finished before step F can begin.\n	Step A must be finished before step B can begin.\nStep A must be finished before step D can begin.\nStep B must be finished before step E can begin.\nStep D must be finished before step E can begin.\nStep F must be finished before step E can begin."

	if result := BuildOrderParallel(in, 2, 0); result != 15 {
		t.Errorf("Incorrect result %v, expected: %v", result, 15)
	}

}

func Test_PartTwo(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day7/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	in := string(dat)

	if result := BuildOrderParallel(in, 5, 60); result != 15 {
		t.Errorf("Incorrect result %v, expected: %v", result, 15)
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
	graph := createGraph(in)

	queue := createInitialqueue(graph)
	for _, n := range queue {
		built[n] = true
	}

	for len(queue) > 0 {
		node := graph[queue[0]]
		queue = queue[1:]
		allDependencies := true
		for _, d := range node.DependsOn {
			if _, ok := built[d]; !ok {
				allDependencies = false
			}
		}

		if allDependencies {
			out += node.Name
			built[node.Name] = true
			for _, i := range node.Builds {
				if _, ok := inLine[i]; !ok {
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

func createGraph(in string) map[string]*Node {
	graph := make(map[string]*Node)
	lines := strings.Split(in, "\n")
	for _, l := range lines {
		n, d := ProcessLine(l)
		if node, ok := graph[n]; !ok {
			graph[n] = &Node{n, []string{d}, []string{}}
		} else {
			node.DependsOn = append(graph[n].DependsOn, d)
		}
		if node, ok := graph[d]; !ok {
			graph[d] = &Node{d, []string{}, []string{n}}
		} else {
			node.Builds = append(graph[d].Builds, n)
		}
	}
	return graph
}

func BuildOrderParallel(in string, numberOfWorkers, stepDuration int) int {
	duration := 0
	processedNodes := make(map[string]bool)
	built := make(map[string]bool)

	graph := createGraph(in)

	queue := createInitialqueue(graph)

	workers := make([]*Worker, 0)
	for i := 0; i < numberOfWorkers; i++ {
		workers = append(workers, &Worker{ExtraStepDuration: stepDuration})
	}

	for len(graph) > len(built) {
		fmt.Println(fmt.Sprintf("Second %v -----", duration))
		for _, w := range workers {
			if w.IsFree() {
				// Mark as Done
				if w.Node != "" {
					built[w.Node] = true
					node := graph[w.Node]
					for _, i := range node.Builds {
						if _, ok := processedNodes[i]; !ok {
							processedNodes[i] = true
							queue = append(queue, i)
						}
					}
					w.Node = ""
				}
				//Assign new work
				q, node := GetNextWorkableNode(graph, queue, built)
				queue = q
				if node != ""{
					w.AssignWork(node)
					processedNodes[node] = true
				}
			}
			fmt.Println(fmt.Sprintf("Worker Status: %v", w))
			w.Tick()
		}
		duration += 1
	}

	return duration - 1
}

func createInitialqueue(graph map[string]*Node) []string {
	queue := make([]string, 0)
	for k, v := range graph {
		if len(v.DependsOn) == 0 {
			queue = append(queue, k)
		}
	}
	sort.Strings(queue)
	return queue
}

func IsNodeReadyToBuild(node *Node, built map[string]bool) bool {
	for _, d := range node.DependsOn {
		if _, ok := built[d]; !ok {
			return false
		}
	}
	return true
}

func GetNextWorkableNode(graph map[string]*Node, q []string, built map[string]bool) (queue []string, out string) {
	out = ""
	sort.Strings(q)
	if len(q) < 1 {
		return q, out
	}
	for i := 0; i < len(q); i++ {
		n := q[i]
		if IsNodeReadyToBuild(graph[n], built) {
			queue = append(q[:i], q[i+1:]...)
			return queue, n
		}
	}
	return q, out
}
