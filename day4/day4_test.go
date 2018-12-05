package day4

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type Entry struct {
	Year int
	Month int
	Day int
	Hour int
	Minute int
	Text string
}

type ChronologicalSortedEntries []Entry

func (b ChronologicalSortedEntries) Len() int {
	return len(b)
}

func (b ChronologicalSortedEntries) Swap(i, j int) {
		b[i], b[j] = b[j], b[i]
}

func (b ChronologicalSortedEntries) Less(i, j int) bool {
	return b[i].Year < b[j].Year ||
		(b[i].Year ==  b[j].Year &&
			b[i].Month <  b[j].Month) ||
		(b[i].Year ==  b[j].Year &&
			b[i].Month ==  b[j].Month &&
			b[i].Day <  b[j].Day) ||
		(b[i].Year ==  b[j].Year &&
			b[i].Month ==  b[j].Month &&
			b[i].Day ==  b[j].Day &&
			b[i].Hour <  b[j].Hour) ||
		(b[i].Year ==  b[j].Year &&
			b[i].Month ==  b[j].Month &&
			b[i].Day ==  b[j].Day &&
			b[i].Hour == b[j].Hour &&
			b[i].Minute < b[j].Minute)

}

func CreateEntryFromLine(s string) Entry {
	newEntry := Entry{}
	newEntry.Year, _ = strconv.Atoi(s[1:5])
	newEntry.Month, _ = strconv.Atoi(s[6:8])
	newEntry.Day, _ = strconv.Atoi(s[9:11])
	newEntry.Hour, _ = strconv.Atoi(s[12:14])
	newEntry.Minute, _ = strconv.Atoi(s[15:17])
	newEntry.Text = s[19:]
	return newEntry
}

func GetGuardFromText(s string) int {
	tokens := strings.Split(s, " ")
	guardId, _ := strconv.Atoi(tokens[1][1:])
	return guardId
}

func Test_CreateEntryFromLine(t *testing.T){
	testCases := []struct{
		In string
		E Entry
	}{
		{
			"[1518-11-01 00:00] Guard #10 begins shift",
			Entry{1518,11,1,0,0,"Guard #10 begins shift"},
		},
		{
			"[1518-11-01 00:05] falls asleep",
			Entry{1518,11,1,0,5,"falls asleep"},
		},
		{
			"[1518-11-01 00:25] wakes up",
			Entry{1518,11,1,0,25,"wakes up"},
		},
	}

	for _, tc := range testCases {
		if entry := CreateEntryFromLine(tc.In);tc.E != entry {
			t.Errorf("Expected %v, Got %v", tc.E, entry)
		}
	}
}

func Test_PartOneExample(t *testing.T) {
	in := "[1518-11-01 00:00] Guard #10 begins shift\n[1518-11-01 00:05] falls asleep\n[1518-11-01 00:25] wakes up\n[1518-11-01 00:30] falls asleep\n[1518-11-01 00:55] wakes up\n[1518-11-01 23:58] Guard #99 begins shift\n[1518-11-02 00:40] falls asleep\n[1518-11-02 00:50] wakes up\n[1518-11-03 00:05] Guard #10 begins shift\n[1518-11-03 00:24] falls asleep\n[1518-11-03 00:29] wakes up\n[1518-11-04 00:02] Guard #99 begins shift\n[1518-11-04 00:36] falls asleep\n[1518-11-04 00:46] wakes up\n[1518-11-05 00:03] Guard #99 begins shift\n[1518-11-05 00:45] falls asleep\n[1518-11-05 00:55] wakes up"

	if val := PartOne(in); val != 240 {
		t.Errorf("Got %d, expected: %d", val, 240)
	}
}

func Test_PartTwoExample(t *testing.T) {
	in := "[1518-11-01 00:00] Guard #10 begins shift\n[1518-11-01 00:05] falls asleep\n[1518-11-01 00:25] wakes up\n[1518-11-01 00:30] falls asleep\n[1518-11-01 00:55] wakes up\n[1518-11-01 23:58] Guard #99 begins shift\n[1518-11-02 00:40] falls asleep\n[1518-11-02 00:50] wakes up\n[1518-11-03 00:05] Guard #10 begins shift\n[1518-11-03 00:24] falls asleep\n[1518-11-03 00:29] wakes up\n[1518-11-04 00:02] Guard #99 begins shift\n[1518-11-04 00:36] falls asleep\n[1518-11-04 00:46] wakes up\n[1518-11-05 00:03] Guard #99 begins shift\n[1518-11-05 00:45] falls asleep\n[1518-11-05 00:55] wakes up"

	if val := PartTwo(in); val != 4455 {
		t.Errorf("Got %d, expected: %d", val, 4455)
	}
}

func Test_PartOne(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day4/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	input := string(dat)

	if val := PartOne(input); val != 240 {
		t.Errorf("Got %d, expected: %d", val, 240)
	}
}

func Test_PartTwo(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day4/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	input := string(dat)

	if val := PartTwo(input); val != 240 {
		t.Errorf("Got %d, expected: %d", val, 240)
	}
}

func PartOne(s string) int {
	var entries []Entry

	lines := strings.Split(s, "\n")
	for _, l := range lines {
		e := CreateEntryFromLine(l)
		entries = append(entries, e)
	}

	sort.Sort(ChronologicalSortedEntries(entries))

	totalTimeAsleep, minuteAsleepPerGuard := ProcessEntries(entries)

	rankedByAsleepTime := rankByValue(totalTimeAsleep)
	sleepyGuard := rankedByAsleepTime[0].Key
	minutesRankedByAsleepForTheGuard := rankByValue(minuteAsleepPerGuard[sleepyGuard])
	mostAsleepMinuteForGuard := minutesRankedByAsleepForTheGuard[0].Key

	return sleepyGuard * mostAsleepMinuteForGuard
}

func PartTwo(s string) int {
	var entries []Entry

	lines := strings.Split(s, "\n")
	for _, l := range lines {
		e := CreateEntryFromLine(l)
		entries = append(entries, e)
	}

	sort.Sort(ChronologicalSortedEntries(entries))

	_, minuteAsleepPerGuard := ProcessEntries(entries)
	sleepyGuard := 0
	selectedMinute := -1
	selectedMinuteTimes := -1
	for guard, timetable := range minuteAsleepPerGuard {
		for minute, times := range timetable {
			if times > selectedMinuteTimes {
				sleepyGuard = guard
				selectedMinute = minute
				selectedMinuteTimes = times
			}
		}
	}

	return sleepyGuard * selectedMinute
}

func ProcessEntries(entries []Entry) (map[int]int, map[int]map[int]int) {
	totalTimeAsleep := make(map[int]int)
	minuteAsleepPerGuard := make(map[int]map[int]int)
	currentGuard := 0
	beginSleepMinute := 0
	for _, e := range entries {
		if strings.Contains(e.Text, "begins shift") {
			currentGuard = GetGuardFromText(e.Text)
			if _, ok := minuteAsleepPerGuard[currentGuard]; !ok {
				minuteAsleepPerGuard[currentGuard] = make(map[int]int)
			}
		} else if strings.Contains(e.Text, "falls asleep") {
			beginSleepMinute = e.Minute
		} else {
			totalTimeAsleep[currentGuard] += e.Minute - beginSleepMinute
			for minuteAsleep := beginSleepMinute; minuteAsleep < e.Minute ; minuteAsleep++ {
				if _, ok := minuteAsleepPerGuard[currentGuard][minuteAsleep]; !ok {
					minuteAsleepPerGuard[currentGuard][minuteAsleep] = 1
				} else {
					minuteAsleepPerGuard[currentGuard][minuteAsleep] += 1
				}

			}
		}
	}

	return totalTimeAsleep, minuteAsleepPerGuard

}

func rankByValue(timeAsleep map[int]int) PairList{
	pl := make(PairList, len(timeAsleep))
	i := 0
	for k, v := range timeAsleep {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key int
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }