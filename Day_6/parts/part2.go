package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func Part2(file *os.File) {
	s := bufio.NewScanner(file)
	pass := 0
	times := []string{}
	records := []string{}
	for s.Scan() {
		re := regexp.MustCompile(`[0-9]+`)
		if pass == 0 {
			r := re.FindAllString(s.Text(), -1)
			for _, v := range r {
				times = append(times, v)
			}
		} else {
			r := re.FindAllString(s.Text(), -1)
			for _, v := range r {
				records = append(records, v)
			}
		}
		pass++
	}
    fmt.Println(times)

	var e error
	time, e := strconv.Atoi(strings.Join(times, ""))
	record, e := strconv.Atoi(strings.Join(records, ""))
	if e != nil {
		panic(e)
	}

    var wg sync.WaitGroup
    var wins atomic.Uint64

    timeDelta := time
    countStart := 0
    test := 1
    for test <= 92 {
        fmt.Printf("Starting goroutine: %v/%v\n", test, 92)
        timeDelta -= 500000
        if timeDelta > 0 {
            countWins(countStart, countStart + 500000, &time, &record, &wins, &wg)
            countStart += 500000
        } else {
            x := 500000 - int(math.Abs(float64(timeDelta)))
            countWins(countStart, countStart + x, &time, &record, &wins, &wg)
            fmt.Println("CONTROL: ", countStart, countStart+x)
            timeDelta = 0
        }
        test++
    }
    wg.Wait()
    fmt.Println("Part 2 solution: ", wins.Load())
}

func countWins(start int, end int, time *int, record *int, wins *atomic.Uint64, wg *sync.WaitGroup) {
    wg.Add(1)
    go func(start int, end int, time *int, record *int, wins *atomic.Uint64, wg *sync.WaitGroup) {
        for i := start+1; i <= end; i++ {
            tRemaining := *time - i
            distance := tRemaining * i
            if distance > *record {
                wins.Add(1)
            }
        }
        fmt.Println("Goroutine finished!")
        wg.Done()
    }(start, end, time, record, wins, wg)
}
