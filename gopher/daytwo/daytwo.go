package daytwo

import (
	"bufio"
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

// isSafe checks if a report is safe WITHOUT removing any level.
func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	// Determine trend: increasing or decreasing
	increasing := levels[1] > levels[0]
	decreasing := levels[1] < levels[0]

	if !increasing && !decreasing {
		return false // equal numbers are not allowed
	}

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 { // no equal numbers allowed
			return false
		}

		if increasing && diff < 0 {
			return false
		}

		if decreasing && diff > 0 {
			return false
		}

		if diff > 3 || diff < -3 {
			return false
		}
	}

	return true
}

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
		isSafe := isSafeReport(report)
		if isSafe {
			safe_reports++
		}
	}
	fmt.Println("Total safe reports :", safe_reports)
}
func isSafeWithDampener(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	// Try removing each level one by one
	for i := 0; i < len(levels); i++ {
		newLevels := append([]int{}, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)
		if isSafe(newLevels) {
			return true
		}
	}
	return false
}

func PartTwo() {
	file, err := os.Open("../input-two.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Fields(line)
		levels := make([]int, len(parts))
		for i, p := range parts {
			n, _ := strconv.Atoi(p)
			levels[i] = n
		}

		if isSafeWithDampener(levels) {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}
