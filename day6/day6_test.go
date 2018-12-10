package day6

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"testing"
)

type Point struct {
	X,Y int
}

func Test_PartOneExample(t *testing.T) {
	in := "1, 1\n1, 6\n8, 3\n3, 4\n5, 5\n8, 9"

	if out := PartOne(in); out != 17 {
		t.Errorf("Something wrong, got %v", out)
	}
}

func Test_PartTwoExample(t *testing.T) {
	in := "1, 1\n1, 6\n8, 3\n3, 4\n5, 5\n8, 9"

	if out := PartTwo(in, 32); out != 16 {
		t.Errorf("Something wrong, got %v", out)
	}
}

func Test_PartOne(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day6/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	in := string(dat)

	if out := PartOne(in); out != 3420 {
		t.Errorf("Something wrong, got %v", out)
	}
}

func Test_PartTwo(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day6/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	in := string(dat)

	if out := PartTwo(in,10000); out != 46667 {
		t.Errorf("Something wrong, got %v", out)
	}
}

func PartOne(in string) int {
	points := ReadLines(in)
	area := make(map[Point]int)
	areaForPoint := make(map[int]int)
	finitePoints := make([]int,0)

	minX,minY,maxX,maxY := FindLimitsFrom(points)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y ++ {
			p := Point{x,y}
			closest := ClosestTo(p,points)
			area[p] = closest
			if _, ok := areaForPoint[closest]; !ok {
				areaForPoint[closest] = 1
			} else {
				areaForPoint[closest] += 1
			}
		}
	}

	for i, p := range points {
		if AreaIsFinite(p,area,minX,minY,maxX,maxY) {
			finitePoints = append(finitePoints, i)
		}
	}

	max := 0
	for _, p := range finitePoints {
		if areaForPoint[p] > max {
			max = areaForPoint[p]
		}
	}

	return max
}

func ReadLines(in string) []Point {
	points := make([]Point,0)
	lines := strings.Split(in, "\n")
	for _, l := range lines {
		points = append(points,LineToPoint(l))
	}
	return points
}

func LineToPoint(in string) Point {
	p := Point{}
	coords := strings.Split(in, ", ")
	p.X, _ = strconv.Atoi(coords[0])
	p.Y, _ = strconv.Atoi(coords[1])
	return p
}

func FindLimitsFrom(points []Point) (minX,minY,maxX,maxY int) {
	minX = math.MaxInt16
	minY = math.MaxInt16
	maxX = math.MinInt16
	maxY = math.MinInt16
	for _, p := range points {
		minX = int(math.Min(float64(minX),float64(p.X)))
		minY = int(math.Min(float64(minY),float64(p.Y)))
		maxX = int(math.Max(float64(maxX),float64(p.X)))
		maxY = int(math.Max(float64(maxY),float64(p.Y)))
	}
	return
}

func ClosestTo(point Point, points []Point) int {
	minManhattanDistance := math.MaxInt16
	closestTo := 0
	tie := false
	for i, p := range points {
		distance := int(math.Abs(float64(p.X - point.X))+math.Abs(float64(p.Y - point.Y)))
		if minManhattanDistance > distance {
			tie = false
			minManhattanDistance = distance
			closestTo = i
		} else if minManhattanDistance == distance {
			tie = true
			closestTo = -1
		}
	}

	if tie {
		return -1
	}

	return closestTo
}

func AreaIsFinite(p Point, area map[Point]int, minX,minY,maxX,maxY int) bool{
	point := area[Point{p.X, p.Y}]

	for x := minX; x < maxX; x++ {
		if area[Point{x,minY}] == point {
			return false
		}
	}

	for x := minX; x < maxX; x++ {
		if area[Point{x,maxY}] == point {
			return false
		}
	}

	for y := minY; y < maxY; y++ {
		if area[Point{minX,y}] == point {
			return false
		}
	}

	for y := minY; y < maxY; y++ {
		if area[Point{maxX,y}] == point {
			return false
		}
	}


	return true

}

func PartTwo(in string, limit int) int {
	points := ReadLines(in)
	sumArea := 0

	minX,minY,maxX,maxY := FindLimitsFrom(points)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y ++ {
			p := Point{x,y}
			distance := 0
			for _, pc := range points {
				distance += int(math.Abs(float64(p.X - pc.X))+math.Abs(float64(p.Y - pc.Y)))
			}
			if distance < limit {
				sumArea += 1
			}
		}
	}

	return sumArea
}