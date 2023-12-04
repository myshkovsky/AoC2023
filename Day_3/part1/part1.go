package part1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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

	numLines := len(lines)
	maxLineLength := len(lines[0])

	addPart := func(str string) {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		sumParts += num
	}

	re := regexp.MustCompile(`[^.0-9]+`)
	checkUpDown := func(line int, start int, end int, up bool) bool {
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

	checkSides := func(line int, start int, end int) bool {
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

	for k, v := range nMap {
		for _, y := range v {
			if k == 0 {
				if checkSides(k, y[0], y[1]) || checkUpDown(k, y[0], y[1], false) {
					addPart(lines[k][y[0]:y[1]])
					continue
				}
			} else if k == numLines-1 {
				if checkUpDown(k, y[0], y[1], true) || checkSides(k, y[0], y[1]) {
					addPart(lines[k][y[0]:y[1]])
					continue
				}
			} else {
				if checkSides(k, y[0], y[1]) ||
					checkUpDown(k, y[0], y[1], true) ||
					checkUpDown(k, y[0], y[1], false) {
					addPart(lines[k][y[0]:y[1]])
					continue
				}
			}
		}
	}
	fmt.Println("Part 1 solution: ", sumParts)
}
