package daytwo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPartsOfLine(line string) []int {
	parts := strings.Fields(line)
	levels := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println(err)
			continue
		}
		levels[i] = num
	}
	return levels
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafeReport(levels []int) bool {
	var direction int // +1 for increasing, -1 for decreasing
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff == 0 || abs(diff) > 3 { // must be nonzero and <= 3
			return false
		}
		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else {
				direction = -1
			}
		} else {
			if direction == 1 && diff < 0 {
				return false
			}
			if direction == -1 && diff > 0 {
				return false
			}
		}
	}
	return true
}
func sortArray(arr []int, desc bool) {
	n := len(arr)
	if desc {
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	} else {
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
}

// func compareArray(arr1 []int, arr2 []int) bool {
// 	if len(arr1) != len(arr2) {
// 		return false
// 	}
// 	for i := 0; i < len(arr1); i++ {
// 		if arr1[i] != arr2[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func PartOne() {
	bytes, err := os.ReadFile("../input-two.txt")
	if err != nil {
		panic(err)
	}

	data := string(bytes)

	lines := strings.Split(data, "\n")

	reports := make([][]int, len(lines))

	for i, line := range lines {
		reports[i] = getPartsOfLine(line)
	}

	var safe_reports int = 0
	for _, report := range reports {
		if isSafeReport(report) {
			safe_reports++
		}
	}
	fmt.Println("Total safe reports :", safe_reports)
}

func isSafeReportWithDampener(level []int) bool {
	return true
}

func PartTwo() {
	bytes, err := os.ReadFile("../input-two.txt")
	if err != nil {
		panic(err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")
	reports := make([][]int, len(lines))

	for i, line := range lines {
		reports[i] = getPartsOfLine(line)
	}

	var safe_reports int
	for _, report := range reports {
		if isSafeReportWithDampener(report) {
			safe_reports++
		}
	}

	fmt.Println("safe reports with dampener :", safe_reports)
}
