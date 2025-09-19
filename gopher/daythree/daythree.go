package daythree

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func PartOne() {
	file, err := os.Open("../input-three.txt")
	if err != nil {
		panic(err)
	}

	var scanner *bufio.Scanner = bufio.NewScanner(file)
	var input string = ""

	for scanner.Scan() {
		input += scanner.Text()
	}

	sum := 0
	n := len(input)

	for i := 0; i < n; i++ {
		if i+4 <= n && input[i:i+4] == "mul(" {
			j := i + 4
			start := j
			for j < n && unicode.IsDigit(rune(input[j])) {
				j++
			}
			if j == start || j-start > 3 {
				continue
			}
			xStr := input[start:j]
			if j >= n || input[j] != ',' {
				continue
			}
			j++
			start = j
			for j < n && unicode.IsDigit(rune(input[j])) {
				j++
			}
			if j == start || j-start > 3 {
				continue
			}
			yStr := input[start:j]
			if j >= n || input[j] != ')' {
				continue
			}
			x, _ := strconv.Atoi(xStr)
			y, _ := strconv.Atoi(yStr)
			sum += x * y
		}
	}

	fmt.Println(sum)
}

func PartTwo() {
	file, err := os.Open("../input-three.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}
	sum := 0
	enabled := true
	n := len(input)
	for i := 0; i < n; i++ {
		if i+7 <= n && input[i:i+7] == "don't()" {
			enabled = false
			i += 6
			continue
		}
		if i+4 <= n && input[i:i+4] == "do()" {
			enabled = true
			i += 3
			continue
		}
		if i+4 <= n && input[i:i+4] == "mul(" {
			j := i + 4
			start := j
			for j < n && unicode.IsDigit(rune(input[j])) {
				j++
			}
			if j == start || j-start > 3 {
				continue
			}
			xStr := input[start:j]
			if j >= n || input[j] != ',' {
				continue
			}
			j++

			start = j
			for j < n && unicode.IsDigit(rune(input[j])) {
				j++
			}
			if j == start || j-start > 3 {
				continue
			}
			yStr := input[start:j]
			if j >= n || input[j] != ')' {
				continue
			}
			if enabled {
				x, _ := strconv.Atoi(xStr)
				y, _ := strconv.Atoi(yStr)
				sum += x * y
			}
		}
	}
	fmt.Println(sum)
}
