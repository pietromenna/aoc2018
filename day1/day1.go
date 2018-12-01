package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SolvePartOne(filePath string) (sum int64) {
	dat, err := ioutil.ReadFile(filePath)
	check(err)

	inputInString := string(dat)
	tokens := strings.Split(inputInString, "\n")

	sum = 0
	for _, token := range tokens {
		i, _ := strconv.Atoi(token)
		sum += int64(i)
	}

	fmt.Println(fmt.Sprintf("END FREQ %d", sum))
	return sum
}

func SolvePartTwo(filePath string) {
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	seenFrequencies := make(map[int64]int64)

	inputInString := string(dat)
	tokens := strings.Split(inputInString, "\n")

	var sum int64 = 0
	seenFrequencies[0] = 1
	found := false
	for !found {
		for _, token := range tokens {
			i, _ := strconv.Atoi(token)
			sum += int64(i)

			if _, ok := seenFrequencies[sum]; !ok {
				seenFrequencies[sum] = 1
			} else {
				found = true
				fmt.Println(fmt.Sprintf("DUPLICATE AT%d",sum))
				return
			}
		}
	}

}
