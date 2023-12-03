package part2

import (
	"AoC2023/solutions/day1/part1"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func replaceInner(str string, replacement int, startIndex int, endIndex int) string{
    return str[:startIndex+1] + strconv.Itoa(replacement) + str[endIndex-1:] 
}

func Solve(file *os.File) {
    scanner := bufio.NewScanner(file)
    digitWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    digitsArray := []int{}
    for scanner.Scan() {
        text := scanner.Text()
        for i := 0; i < len(digitWords); i++ {
            r := regexp.MustCompile(fmt.Sprintf("(%v)", digitWords[i]))
            for r.FindStringIndex(text) != nil {
                index := r.FindStringIndex(text)
                text = replaceInner(text, i+1, index[0], index[1])
            }
        }
        r := regexp.MustCompile("[0-9]+").FindAllString(text, -1)
        part1.IsolateDigits(&digitsArray, r)
    }
    sum := 0
	for _, v := range digitsArray {
		sum += v
	}
    fmt.Println("Total sum: ", sum)
}
