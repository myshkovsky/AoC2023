package part1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func CheckUpDown(line int, start int, end int, up bool, lines []string, re *regexp.Regexp, maxLineLength int) bool {
	var direction int
	if up {
		direction = -1
	} else {
		direction = 1
	}
	var r []int
	if start == 0 {
		r = re.FindStringIndex(lines[line+direction][start : end+1])
	} else if end == maxLineLength {
		r = re.FindStringIndex(lines[line+direction][start-1 : end])
	} else {
		r = re.FindStringIndex(lines[line+direction][start-1 : end+1])
	}
	return r != nil
}

func CheckSides(line int, start int, end int, lines []string, re *regexp.Regexp, maxLineLength int) bool {
	isLeftSymbol := func() bool {
		return string(lines[line][start-1]) != "."
	}
	isRightSymbol := func() bool {
		return string(lines[line][end]) != "."
	}
	if start == 0 {
		return isRightSymbol()
	} else if end == maxLineLength {
		return isLeftSymbol()
	} else {
		return isRightSymbol() || isLeftSymbol()
	}
}

func Solve(file *os.File) {
	sumParts := 0
	s := bufio.NewScanner(file)
	lines := []string{}
	currentLine := 0
	nMap := make(map[int][][]int)

	for s.Scan() {
		lines = append(lines, s.Text())
		r := regexp.MustCompile(`[0-9]+`).FindAllStringIndex(s.Text(), -1)
		if r != nil {
			nMap[currentLine] = r
		}
		currentLine++
	}

	addPart := func(str string) {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		sumParts += num
	}

	numLines := len(lines)
	maxLineLength := len(lines[0])
	re := regexp.MustCompile(`[^.0-9]+`)

	for k, v := range nMap {
		for _, y := range v {
			if k == 0 {
				if CheckSides(k, y[0], y[1], lines, re, maxLineLength) ||
					CheckUpDown(k, y[0], y[1], false, lines, re, maxLineLength) {
					addPart(lines[k][y[0]:y[1]])
					continue
				}
			} else if k == numLines-1 {
				if CheckUpDown(k, y[0], y[1], true, lines, re, maxLineLength) ||
					CheckSides(k, y[0], y[1], lines, re, maxLineLength) {
					addPart(lines[k][y[0]:y[1]])
					continue
				}
			} else {
				if CheckSides(k, y[0], y[1], lines, re, maxLineLength) ||
					CheckUpDown(k, y[0], y[1], true, lines, re, maxLineLength) ||
					CheckUpDown(k, y[0], y[1], false, lines, re, maxLineLength) {
					addPart(lines[k][y[0]:y[1]])
					continue
				}
			}
		}
	}
	fmt.Println("Part 1 solution: ", sumParts)
}
