package part1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func determineScore(str string) int {
	r := regexp.MustCompile(`[0-9]+`).FindString(str)
	n, e := strconv.Atoi(r)
	if e != nil {
		panic(e)
	}
	return n
}

func determineColor(str string) string {
	if strings.Contains(str, "red") {
		return "red"
	} else if strings.Contains(str, "green") {
		return "green"
	} else {
		return "blue"
	}
}

func GetGameId(s *bufio.Scanner) int {
	idstr := regexp.MustCompile("[0-9]{1,3}").FindString(s.Text())
	id, e := strconv.Atoi(idstr)
	if e != nil {
		panic(e)
	}
	return id
}

func GetHighestRGB(s *bufio.Scanner) map[string]int{
	r := regexp.MustCompile(`([0-9]+\s(red|green|blue))`).FindAllString(s.Text(), -1)
	m := make(map[string]int)
	for _, v := range r {
		score := determineScore(v)
		color := determineColor(v)
		if m[color] < score {
			m[color] = score
		}
	}
    return m
}

func Solve(file *os.File) {
	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)
	games := 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		id := GetGameId(s)
		m := GetHighestRGB(s)
		if m["red"] <= maxRed && m["green"] <= maxGreen && m["blue"] <= maxBlue {
			games += id
		}
	}
	fmt.Println("Sum of valid game IDs: ", games)
}
