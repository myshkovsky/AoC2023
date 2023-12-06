package solutions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1(file *os.File) {
	s := bufio.NewScanner(file)
    times := []int{}
    records := []int{}
	re := regexp.MustCompile(`[0-9]+`)
    pass := 0
    for s.Scan() {
		if pass == 0 {
			r := re.FindAllString(s.Text(), -1)
			for _, v := range r {
				n, e := strconv.Atoi(v)
				if e != nil {
					panic(e)
				}
				times = append(times, n)
			}
		} else {
			r := re.FindAllString(s.Text(), -1)
			for _, v := range r {
				n, e := strconv.Atoi(v)
				if e != nil {
					panic(e)
				}
				records = append(records, n)
			}
		}
        pass++
	}
    waysToWin := []int{}
	for k, time := range times {
		record := records[k]
		wins := 0
		for i := 0; i <= time; i++ {
			timeRemaining := time - i
			distance := timeRemaining * i
			if distance > record {
				wins++
			}
		}
		waysToWin = append(waysToWin, wins)
	}
    sum := 1
    for _, v := range waysToWin {
        sum *= v
    }
    fmt.Println(sum)
}
