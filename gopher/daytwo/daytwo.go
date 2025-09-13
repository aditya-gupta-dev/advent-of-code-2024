package daytwo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPartsOfLine(line string) []int {
	parts := strings.Fields(line)
	levels := make([]int, 0, len(parts))
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

func PartOne() {
	bytes, err := os.ReadFile("../input-two.txt")
	if err != nil {
		panic(err)
	}

	data := string(bytes)

	lines := strings.Split(data, "\n")

	reports := make([][]int, 0)

	for i, line := range lines {
		reports[i] = getPartsOfLine(line)
	}
	fmt.Println(reports[1])
}
