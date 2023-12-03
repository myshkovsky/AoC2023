package part1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func IsolateDigits(arr *[]int, s []string) {
	digits := strings.Join(s, "")
	numStr := fmt.Sprintf("%v%v", digits[:1], digits[len(digits)-1:])
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	*arr = append(*arr, num)
}

func Solve(file *os.File) {
	scanner := bufio.NewScanner(file)
	digitsArray := []int{}
	for scanner.Scan() {
        r := regexp.MustCompile("[0-9]+").FindAllString(scanner.Text(), -1)
        IsolateDigits(&digitsArray, r)
	}
    sum := 0
	for _, v := range digitsArray {
		sum += v
	}
	fmt.Println("Total sum: ", sum)
}
