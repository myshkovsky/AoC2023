package solution

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetWinningNumbers(str string) map[int]bool {
	m := make(map[int]bool)
	r := regexp.MustCompile(`:\s+((\d)+(\s)*)+`).FindString(str)
	r = strings.Trim(strings.Replace(r, "  ", " ", -1), ": ")
	for _, v := range strings.Split(r, " ") {
		n, e := strconv.Atoi(v)
		if e != nil {
			panic(e)
		}
		m[n] = true
	}
	return m
}

func GetYourNumbers(str string) []int {
	r := regexp.MustCompile(`\|\s+((\d)+(\s)*)+`).FindString(str)
	r = strings.Trim(strings.Replace(r, "  ", " ", -1), "| ")
    var nums []int
	for _, v := range strings.Split(r, " ") {
		n, e := strconv.Atoi(v)
		if e != nil {
			panic(e)
		}
        nums = append(nums, n)
	}
    return nums
}

func Part1(file *os.File) {
	sum := 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		m := GetWinningNumbers(s.Text())
        nums := GetYourNumbers(s.Text())
		hits := 0
		for _, v := range nums {
			if m[v] {
				hits++
			}
		}
		if hits > 2 {
			sum += int(math.Pow(2, float64(hits-1)))
		} else {
			sum += hits
		}
	}
	fmt.Println("Part 1 solution: ", sum)
}
